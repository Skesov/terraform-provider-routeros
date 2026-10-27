package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"
)

// Menus that are not configuration: runtime tables, logs, files and tools
// whose print output is large or changes on every read.
var skipPrefixes = map[string]string{
	"/file":                              "files, not configuration",
	"/log":                               "log buffer",
	"/system/history":                    "undo history",
	"/ip/firewall/connection":            "connection tracking table",
	"/ipv6/firewall/connection":          "connection tracking table",
	"/tool/sniffer":                      "packet capture tool",
	"/tool/torch":                        "traffic monitor tool",
	"/console":                           "console internals",
	"/terminal":                          "console internals",
	"/environment":                       "script variables",
	"/task":                              "console internals",
	"/safe-mode":                         "console internals",
	"/special-login":                     "console internals",
	"/system/script/job":                 "running scripts",
	"/interface/wifi/registration-table": "connected clients",
}

// Arguments that are command options, not properties.
var commandArgs = map[string]bool{"numbers": true, "copy-from": true, "place-before": true}

// Arg is a command argument: a property that can be set.
type Arg struct {
	Description string   `json:"description,omitempty"`
	Syntax      string   `json:"syntax,omitempty"`
	Values      []string `json:"values,omitempty"`
}

type Menu struct {
	Kind    string            `json:"kind"` // singleton, list, fixed-list, none
	Add     map[string]*Arg   `json:"add,omitempty"`
	Set     map[string]*Arg   `json:"set,omitempty"`
	Print   map[string]string `json:"print,omitempty"` // property -> inferred type
	Rows    int               `json:"rows,omitempty"`
	Static  int               `json:"static_rows,omitempty"` // rows that are neither dynamic nor default
	Error   string            `json:"error,omitempty"`
	Skipped string            `json:"skipped,omitempty"`
}

type Dump struct {
	Version   string           `json:"routeros_version"`
	Board     string           `json:"board"`
	Generated string           `json:"generated"`
	Menus     map[string]*Menu `json:"menus"`
}

type client struct {
	base, user, pass string
	http             *http.Client

	// Names of objects seen on the router, removed from the dump.
	mu    sync.Mutex
	names map[string]bool
}

var nameProps = []string{"name", "comment", "default-name", "user", "ssid", "identity", "address-list", "list"}

func (c *client) addNames(r map[string]any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, k := range nameProps {
		if v, ok := r[k].(string); ok && v != "" {
			c.names[v] = true
		}
	}
}

var reToken = regexp.MustCompile(`[^\s|;:,]+`)

// redact drops router object names from argument values and syntax.
func (c *client) redact(d *Dump) {
	for _, m := range d.Menus {
		for _, args := range []map[string]*Arg{m.Add, m.Set} {
			for _, a := range args {
				vals := a.Values[:0]
				for _, v := range a.Values {
					if !c.names[v] {
						vals = append(vals, v)
					}
				}
				a.Values = vals
				a.Syntax = reToken.ReplaceAllStringFunc(a.Syntax, func(t string) string {
					if c.names[t] {
						return "<name>"
					}
					return t
				})
			}
		}
	}
}

func (c *client) do(method, path string, body any, out any) error {
	var rd io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		rd = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, c.base+"/rest"+path, rd)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.user, c.pass)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %s: %d %s", method, path, resp.StatusCode, strings.TrimSpace(string(data)))
	}
	return json.Unmarshal(data, out)
}

type node struct {
	Name     string `json:"name"`
	NodeType string `json:"node-type"`
	Type     string `json:"type"`
}

type syntax struct {
	Symbol     string `json:"symbol"`
	SymbolType string `json:"symbol-type"`
	Text       string `json:"text"`
}

// inspect is the only POST this tool sends; console/inspect is read-only.
func (c *client) children(path []string) ([]node, error) {
	var res []node
	err := c.do("POST", "/console/inspect", map[string]string{"request": "child", "path": strings.Join(path, ",")}, &res)
	var out []node
	for _, n := range res {
		if n.Type == "child" {
			out = append(out, n)
		}
	}
	return out, err
}

func (c *client) inspect(request string, path []string) ([]map[string]string, error) {
	var res []map[string]string
	err := c.do("POST", "/console/inspect", map[string]string{"request": request, "path": strings.Join(path, ",")}, &res)
	return res, err
}

func (c *client) args(path []string) (map[string]*Arg, error) {
	nodes, err := c.children(path)
	if err != nil {
		return nil, err
	}
	res := map[string]*Arg{}
	for _, n := range nodes {
		if n.NodeType == "arg" && !commandArgs[n.Name] {
			res[n.Name] = &Arg{}
		}
	}
	if syn, err := c.inspect("syntax", path); err == nil {
		for _, s := range syn {
			if a, ok := res[s["symbol"]]; ok && s["symbol-type"] == "explanation" {
				a.Description = s["text"]
			}
		}
	}
	// Syntax and completion can name objects on the router (interfaces,
	// scripts, comments). Only type definitions (symbols starting with an
	// upper-case letter) and completions of pure enums are kept, so the dump
	// holds no router-specific data.
	for name, a := range res {
		argPath := append(append([]string{}, path...), name)
		rawSyntax := false
		if syn, err := c.inspect("syntax", argPath); err == nil {
			var texts []string
			for _, s := range syn {
				if s["symbol-type"] != "definition" || s["text"] == "" {
					continue
				}
				rawSyntax = true
				if sym := s["symbol"]; sym != "" && unicode.IsUpper(rune(sym[0])) {
					texts = append(texts, sym+": "+s["text"])
				}
			}
			a.Syntax = strings.Join(texts, "; ")
		}
		if comp, err := c.inspect("completion", argPath); err == nil {
			for _, s := range comp {
				if s["show"] != "true" {
					continue
				}
				if s["style"] == "arg" || (s["style"] == "none" && !rawSyntax && s["text"] == "") {
					a.Values = append(a.Values, s["completion"])
				}
			}
			sort.Strings(a.Values)
		}
	}
	return res, nil
}

var reInt = regexp.MustCompile(`^-?\d+$`)

func inferType(v string) string {
	switch {
	case v == "true" || v == "false":
		return "bool"
	case reInt.MatchString(v):
		return "int"
	default:
		return "string"
	}
}

func (c *client) menu(path []string, cmds map[string]bool) *Menu {
	m := &Menu{Kind: "none"}
	var err error
	if cmds["add"] {
		m.Kind = "list"
		if m.Add, err = c.args(append(path, "add")); err != nil {
			m.Error = err.Error()
		}
	}
	if cmds["set"] {
		nodes, _ := c.children(append(path, "set"))
		hasNumbers := false
		for _, n := range nodes {
			if n.Name == "numbers" {
				hasNumbers = true
			}
		}
		if m.Kind == "none" {
			m.Kind = "singleton"
			if hasNumbers {
				m.Kind = "fixed-list"
			}
		}
		if m.Set, err = c.args(append(path, "set")); err != nil {
			m.Error = err.Error()
		}
	}
	if m.Kind == "none" || !cmds["print"] {
		return m
	}

	var raw json.RawMessage
	if err := c.do("GET", "/"+strings.Join(path, "/"), nil, &raw); err != nil {
		m.Error = err.Error()
		return m
	}
	var rows []map[string]any
	if len(raw) > 0 && raw[0] == '{' {
		var one map[string]any
		_ = json.Unmarshal(raw, &one)
		rows = []map[string]any{one}
	} else {
		_ = json.Unmarshal(raw, &rows)
	}
	m.Print = map[string]string{}
	for _, r := range rows {
		if r["dynamic"] != "true" && r["default"] != "true" && r["builtin"] != "true" {
			m.Static++
		}
		c.addNames(r)
		for k, v := range r {
			if strings.HasPrefix(k, ".") {
				continue
			}
			// Only the type is kept: values can hold secrets.
			t := inferType(fmt.Sprint(v))
			if old, ok := m.Print[k]; ok && old != t {
				t = "string"
			}
			m.Print[k] = t
		}
	}
	if m.Kind != "singleton" {
		m.Rows = len(rows)
	} else {
		m.Static = 0
	}
	return m
}

func (c *client) walk(path []string, d *Dump, mu *sync.Mutex, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	p := "/" + strings.Join(path, "/")
	for prefix, reason := range skipPrefixes {
		if p == prefix || strings.HasPrefix(p, prefix+"/") {
			if p == prefix {
				mu.Lock()
				d.Menus[p] = &Menu{Kind: "none", Skipped: reason}
				mu.Unlock()
			}
			return
		}
	}

	sem <- struct{}{}
	nodes, err := c.children(path)
	<-sem
	if err != nil {
		mu.Lock()
		d.Menus[p] = &Menu{Kind: "none", Error: err.Error()}
		mu.Unlock()
		return
	}

	cmds := map[string]bool{}
	for _, n := range nodes {
		switch n.NodeType {
		case "cmd":
			cmds[n.Name] = true
		case "dir", "path":
			wg.Add(1)
			go c.walk(append(append([]string{}, path...), n.Name), d, mu, wg, sem)
		}
	}

	if len(path) > 0 && (cmds["add"] || cmds["set"]) {
		sem <- struct{}{}
		m := c.menu(path, cmds)
		<-sem
		mu.Lock()
		d.Menus[p] = m
		mu.Unlock()
	}
}

func runDump(url, out string) error {
	c := &client{
		base:  strings.TrimRight(url, "/"),
		user:  os.Getenv("ROS_USERNAME"),
		pass:  os.Getenv("ROS_PASSWORD"),
		names: map[string]bool{},
		http: &http.Client{
			Timeout:   60 * time.Second,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: os.Getenv("ROS_INSECURE") == "true"}},
		},
	}
	if c.user == "" {
		return fmt.Errorf("ROS_USERNAME and ROS_PASSWORD must be set")
	}

	var res map[string]any
	if err := c.do("GET", "/system/resource", nil, &res); err != nil {
		return err
	}
	d := &Dump{
		Version:   strings.Fields(fmt.Sprint(res["version"]))[0],
		Board:     fmt.Sprint(res["board-name"]),
		Generated: time.Now().UTC().Format("2006-01-02"),
		Menus:     map[string]*Menu{},
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, 8)
	wg.Add(1)
	c.walk(nil, d, &mu, &wg, sem)
	wg.Wait()
	c.redact(d)

	if out == "" {
		out = fmt.Sprintf("schema-%s-%s.json", d.Version, boardSlug(d.Board))
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", " ")
	err := enc.Encode(d)
	b := bytes.TrimRight(buf.Bytes(), "\n")
	if err != nil {
		return err
	}
	keys := make([]string, 0, len(d.Menus))
	for k := range d.Menus {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Fprintf(os.Stderr, "%s %s: %d menus -> %s\n", d.Board, d.Version, len(keys), out)
	return os.WriteFile(out, append(b, '\n'), 0o644)
}

func boardSlug(b string) string {
	b = strings.ToLower(b)
	b = strings.NewReplacer("^", "", " ", "-", "+", "plus").Replace(b)
	return b
}
