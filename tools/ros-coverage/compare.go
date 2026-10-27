package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-routeros/terraform-provider-routeros/routeros"
)

// Resources for other switch chips that share a menu path with the generic
// resource. Their schema does not apply to the dumped boards.
var otherHardware = map[string]string{
	"routeros_interface_ethernet_switch_crs":      "CRS1xx/2xx switch chips",
	"routeros_interface_ethernet_switch_crs_vlan": "CRS1xx/2xx switch chips",
}

// Schema fields that map to command arguments, not properties; the dump
// leaves those arguments out.
var commandFields = map[string]bool{"place_before": true, "numbers": true, "copy_from": true}

type result struct {
	Path, Kind, Resource                       string
	Static                                     int
	OK                                         int
	MissingRW, MissingRO, ROInProvider, Absent []string
	Skipped                                    []string
}

func (r *result) gaps() int {
	return len(r.MissingRW) + len(r.MissingRO) + len(r.ROInProvider) + len(r.Absent)
}

// merge combines dumps from several routers: a menu or property present on
// any of them counts.
func merge(files []string) (*Dump, error) {
	all := &Dump{Menus: map[string]*Menu{}}
	var boards []string
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		var d Dump
		if err := json.Unmarshal(b, &d); err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}
		if all.Version == "" {
			all.Version = d.Version
		}
		boards = append(boards, d.Board+" "+d.Version)
		for p, m := range d.Menus {
			cur, ok := all.Menus[p]
			if !ok {
				cur = &Menu{Kind: m.Kind, Add: map[string]*Arg{}, Set: map[string]*Arg{}, Print: map[string]string{}, Skipped: m.Skipped}
				all.Menus[p] = cur
			}
			for k, v := range m.Add {
				cur.Add[k] = v
			}
			for k, v := range m.Set {
				cur.Set[k] = v
			}
			for k, v := range m.Print {
				cur.Print[k] = v
			}
			cur.Static = max(cur.Static, m.Static)
			cur.Rows = max(cur.Rows, m.Rows)
		}
	}
	all.Board = strings.Join(boards, ", ")
	return all, nil
}

func fieldExists(s map[string]*schema.Schema, field string) (*schema.Schema, bool) {
	top, sub, nested := strings.Cut(field, ".")
	f, ok := s[top]
	if !ok || !nested {
		return f, ok
	}
	if r, ok := f.Elem.(*schema.Resource); ok {
		sf, ok := r.Schema[sub]
		return sf, ok
	}
	return f, true
}

func compareMenu(path string, m *Menu, name string, s map[string]*schema.Schema, ros string) *result {
	r := &result{Path: path, Kind: m.Kind, Resource: name, Static: m.Static}
	configurable := map[string]bool{}
	props := map[string]bool{}
	for k := range m.Set {
		configurable[k], props[k] = true, true
	}
	for k := range m.Add {
		configurable[k], props[k] = true, true
	}
	for k := range m.Print {
		props[k] = true
	}

	seen := map[string]bool{}
	for p := range props {
		field, skipped := routeros.CoverageField(s, ros, p)
		top, _, _ := strings.Cut(field, ".")
		seen[top] = true
		if skipped {
			r.Skipped = append(r.Skipped, p)
			continue
		}
		f, ok := fieldExists(s, field)
		switch {
		case !ok && configurable[p]:
			r.MissingRW = append(r.MissingRW, p)
		case !ok:
			r.MissingRO = append(r.MissingRO, p)
		case configurable[p] && !f.Optional && !f.Required:
			r.ROInProvider = append(r.ROInProvider, p)
		default:
			r.OK++
		}
	}

	for field, f := range s {
		if strings.HasPrefix(field, "___") || seen[field] || commandFields[field] {
			continue
		}
		p := routeros.CoverageProperty(s, ros, field)
		top, _, _ := strings.Cut(p, ".")
		if props[top] || hasPrefix(props, top+".") {
			continue
		}
		// A read-only field missing from print may just be unset on this router.
		if !f.Optional && !f.Required {
			continue
		}
		r.Absent = append(r.Absent, field)
	}
	for _, l := range [][]string{r.MissingRW, r.MissingRO, r.ROInProvider, r.Absent, r.Skipped} {
		sort.Strings(l)
	}
	return r
}

func hasPrefix(m map[string]bool, prefix string) bool {
	for k := range m {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}
	return false
}

func runCompare(files []string, filter string, configured, details bool) error {
	if len(files) == 0 {
		return fmt.Errorf("no schema files given")
	}
	d, err := merge(files)
	if err != nil {
		return err
	}

	byPath := map[string][]string{}
	resources := routeros.Provider().ResourcesMap
	for name, res := range resources {
		p, ok := res.Schema[routeros.MetaResourcePath]
		if !ok || res.DeprecationMessage != "" {
			continue
		}
		byPath[p.Default.(string)] = append(byPath[p.Default.(string)], name)
	}

	paths := make([]string, 0, len(d.Menus))
	for p := range d.Menus {
		paths = append(paths, p)
	}
	sort.Strings(paths)

	var results []*result
	for _, p := range paths {
		m := d.Menus[p]
		if m.Kind == "none" || (filter != "" && !strings.Contains(p, filter)) {
			continue
		}
		if configured && m.Kind != "singleton" && m.Static == 0 {
			continue
		}
		names := byPath[p]
		if len(names) == 0 {
			results = append(results, &result{Path: p, Kind: m.Kind, Static: m.Static})
			continue
		}
		sort.Strings(names)
		for _, n := range names {
			if hw, ok := otherHardware[n]; ok {
				fmt.Fprintf(os.Stderr, "skipped %s: %s only\n", n, hw)
				continue
			}
			results = append(results, compareMenu(p, m, n, resources[n].Schema, d.Version))
		}
	}

	fmt.Printf("RouterOS %s (%s)\n\n", d.Version, d.Board)
	fmt.Println("| Path | Kind | Static rows | Resource | OK | Missing RW | Missing RO | RO in provider | Absent on router |")
	fmt.Println("| --- | --- | --- | --- | --- | --- | --- | --- | --- |")
	var noRes, withGaps, clean int
	for _, r := range results {
		if r.Resource == "" {
			noRes++
			fmt.Printf("| %s | %s | %d | **none** | | | | | |\n", r.Path, r.Kind, r.Static)
			continue
		}
		if r.gaps() == 0 {
			clean++
		} else {
			withGaps++
		}
		fmt.Printf("| %s | %s | %d | %s | %d | %d | %d | %d | %d |\n", r.Path, r.Kind, r.Static, r.Resource,
			r.OK, len(r.MissingRW), len(r.MissingRO), len(r.ROInProvider), len(r.Absent))
	}
	fmt.Printf("\nResources without gaps: %d, with gaps: %d. Menus without a resource: %d.\n", clean, withGaps, noRes)

	if details {
		fmt.Println()
		for _, r := range results {
			if r.Resource == "" || r.gaps() == 0 {
				continue
			}
			fmt.Printf("%s (%s)\n", r.Path, r.Resource)
			for _, x := range []struct {
				t string
				l []string
			}{
				{"missing configurable", r.MissingRW},
				{"missing read-only", r.MissingRO},
				{"read-only in provider", r.ROInProvider},
				{"absent on router", r.Absent},
			} {
				if len(x.l) > 0 {
					fmt.Printf("  %s: %s\n", x.t, strings.Join(x.l, " "))
				}
			}
		}
	}
	return nil
}
