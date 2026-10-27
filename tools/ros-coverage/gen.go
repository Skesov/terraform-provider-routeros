package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/terraform-routeros/terraform-provider-routeros/routeros"
)

// runGen prints draft schema entries for the properties a resource lacks.
// The output is a starting point: review types and descriptions before use.
func runGen(files []string, resource string) error {
	d, err := merge(files)
	if err != nil {
		return err
	}
	res, ok := routeros.Provider().ResourcesMap[resource]
	if !ok {
		return fmt.Errorf("unknown resource %s", resource)
	}
	path := res.Schema[routeros.MetaResourcePath].Default.(string)
	m, ok := d.Menus[path]
	if !ok {
		return fmt.Errorf("%s (%s) is not in the dump", resource, path)
	}
	r := compareMenu(path, m, resource, res.Schema, d.Version)
	for _, p := range r.MissingRW {
		fmt.Print(fieldSnippet(p, argOf(m, p), m.Print[p], true))
	}
	for _, p := range r.MissingRO {
		fmt.Print(fieldSnippet(p, nil, m.Print[p], false))
	}
	return nil
}

// funcName turns routeros_ip_proxy into ResourceIpProxy.
func funcName(resource string) string {
	var b strings.Builder
	b.WriteString("Resource")
	for _, w := range strings.Split(strings.TrimPrefix(resource, "routeros_"), "_") {
		if w == "" {
			continue
		}
		r := []rune(w)
		r[0] = unicode.ToUpper(r[0])
		b.WriteString(string(r))
	}
	return b.String()
}

// runGenResource writes a draft resource, example, import script and (for
// settings menus) an acceptance test for a menu without a resource, under
// repo. Review the output: types come from the router's syntax and print
// output, descriptions from its built-in help.
func runGenResource(files []string, repo, path, resource string) error {
	d, err := merge(files)
	if err != nil {
		return err
	}
	m, ok := d.Menus[path]
	if !ok {
		return fmt.Errorf("%s is not in the dump", path)
	}
	args := map[string]*Arg{}
	for k, v := range m.Set {
		args[k] = v
	}
	for k, v := range m.Add {
		args[k] = v
	}
	props := map[string]bool{}
	for k := range args {
		props[k] = true
	}
	for k := range m.Print {
		props[k] = true
	}
	keys := make([]string, 0, len(props))
	for k := range props {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var skip []string
	var b strings.Builder
	for _, p := range keys {
		rw := args[p] != nil
		switch {
		case p == "comment" && rw:
			b.WriteString("\t\tKeyComment: PropCommentRw,\n")
		case p == "disabled" && rw:
			b.WriteString("\t\tKeyDisabled: PropDisabledRw,\n")
		case p == "dynamic" && !rw:
			b.WriteString("\t\tKeyDynamic: PropDynamicRo,\n")
		case p == "invalid" && !rw:
			b.WriteString("\t\tKeyInvalid: PropInvalidRo,\n")
		case p == "inactive" && !rw:
			b.WriteString("\t\tKeyInactive: PropInactiveRo,\n")
		case p == "default" && !rw:
			b.WriteString("\t\tKeyDefault: PropDefaultRo,\n")
		case p == "name" && m.Kind == "fixed-list":
			// Rows are found by name; the name itself is not written.
			b.WriteString("\t\tKeyName: PropName(\"Name of the item to configure.\"),\n")
			if !rw {
				skip = append(skip, `"name"`)
			}
		case p == "name" && rw && m.Kind == "list":
			desc := "Name of the item."
			if a := args[p]; a.Description != "" {
				desc = sentence(a.Description)
			}
			fmt.Fprintf(&b, "\t\tKeyName: PropName(%q),\n", desc)
		case strings.Contains(p, ".") || strings.HasPrefix(p, "."):
			skip = append(skip, fmt.Sprintf("%q", routeros.KebabToSnake(p)))
		default:
			b.WriteString(fieldSnippet(p, args[p], m.Print[p], rw))
		}
	}

	var create, read, update, del, importer string
	switch m.Kind {
	case "singleton":
		create, read, update, del = "DefaultSystemCreate", "DefaultSystemRead", "DefaultSystemUpdate", "DefaultSystemDelete"
		importer = "schema.ImportStatePassthroughContext"
	case "fixed-list":
		create, read, update, del = "DefaultCreateUpdate", "DefaultRead", "DefaultCreateUpdate", "DefaultSystemDelete"
		importer = "ImportStateCustomContext(resSchema)"
	default:
		create, read, update, del = "DefaultCreate", "DefaultRead", "DefaultUpdate", "DefaultDelete"
		importer = "ImportStateCustomContext(resSchema)"
	}
	skipLine := ""
	if len(skip) > 0 {
		skipLine = "\t\tMetaSkipFields:   PropSkipFields(" + strings.Join(skip, ", ") + "),\n"
	}

	src := fmt.Sprintf(`package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// %[1]s, generated from RouterOS %[8]s by tools/ros-coverage.
func %[2]s() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("%[1]s"),
		MetaId:           PropId(Id),
%[9]s
%[3]s	}

	return &schema.Resource{
		CreateContext: %[4]s(resSchema),
		ReadContext:   %[5]s(resSchema),
		UpdateContext: %[6]s(resSchema),
		DeleteContext: %[7]s(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: %[10]s,
		},

		Schema: resSchema,
	}
}
`, path, funcName(resource), b.String(), create, read, update, del, d.Version, skipLine, importer)

	base := strings.TrimPrefix(resource, "routeros_")
	for _, f := range []string{
		filepath.Join(repo, "routeros", "resource_"+base+".go"),
		filepath.Join(repo, "routeros", "resource_"+base+"_test.go"),
		filepath.Join(repo, "examples", "resources", resource),
	} {
		if _, err := os.Stat(f); err == nil {
			return fmt.Errorf("%s exists; choose another resource name", f)
		}
	}
	if err := os.WriteFile(filepath.Join(repo, "routeros", "resource_"+base+".go"), []byte(src), 0o644); err != nil {
		return err
	}

	exDir := filepath.Join(repo, "examples", "resources", resource)
	if err := os.MkdirAll(exDir, 0o755); err != nil {
		return err
	}
	var example, imp string
	switch m.Kind {
	case "singleton":
		example = fmt.Sprintf("resource %q \"settings\" {\n}\n", resource)
		imp = fmt.Sprintf("terraform import %s.settings .\n", resource)
	case "fixed-list":
		example = fmt.Sprintf("resource %q \"item\" {\n  name = \"<name>\"\n}\n", resource)
		imp = fmt.Sprintf("#The ID can be found via API or the terminal\n#The command for the terminal is -> :put [%s get [print show-ids]]\nterraform import %s.item <name>\n", path, resource)
	default:
		example = fmt.Sprintf("resource %q \"item\" {\n}\n", resource)
		imp = fmt.Sprintf("#The ID can be found via API or the terminal\n#The command for the terminal is -> :put [%s get [print show-ids]]\nterraform import %s.item *1\n", path, resource)
	}
	if err := os.WriteFile(filepath.Join(exDir, "resource.tf"), []byte(example), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(exDir, "import.sh"), []byte(imp), 0o644); err != nil {
		return err
	}

	if m.Kind == "singleton" {
		test := fmt.Sprintf(`package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const test%[1]s = "%[2]s.test"

func TestAcc%[1]sTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, "%[3]s") {
		t.Logf("Test skipped, the minimum required version is %[3]s")
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`+"`"+`%%v

resource "%[2]s" "test" {
}
`+"`"+`, providerConfig),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(test%[1]s),
						),
					},
				},
			})
		})
	}
}
`, strings.TrimPrefix(funcName(resource), "Resource"), resource, d.Version)
		if err := os.WriteFile(filepath.Join(repo, "routeros", "resource_"+base+"_test.go"), []byte(test), 0o644); err != nil {
			return err
		}
	}

	fmt.Printf("\t\t\t%q: %s(),\n", resource, funcName(resource))
	return nil
}

func argOf(m *Menu, p string) *Arg {
	if a, ok := m.Add[p]; ok {
		return a
	}
	return m.Set[p]
}

// A plain integer: no negation, mask, range list or alternative keywords.
var reIntSyntax = regexp.MustCompile(`^\w+: Num; Num: -?\d+\.\.\d+\s+\(integer number\)$`)

func sentence(s string) string {
	s = strings.TrimSpace(strings.ReplaceAll(s, "\r\n", "\n"))
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	s = string(r)
	if !strings.HasSuffix(s, ".") {
		s += "."
	}
	return s
}

func fieldSnippet(prop string, a *Arg, printType string, rw bool) string {
	typ := "TypeString"
	desc := ""
	if a != nil {
		desc = sentence(a.Description)
		isBool := len(a.Values) == 2 && a.Values[0] == "no" && a.Values[1] == "yes"
		switch {
		case isBool:
			typ = "TypeBool"
		case reIntSyntax.MatchString(a.Syntax):
			typ = "TypeInt"
		case len(a.Values) > 0 && len(a.Values) <= 30:
			desc = strings.TrimSpace(desc + " Possible values: `" + strings.Join(a.Values, "`, `") + "`.")
		}
	} else {
		switch printType {
		case "bool":
			typ = "TypeBool"
		case "int":
			typ = "TypeInt"
		}
	}
	name := routeros.KebabToSnake(prop)
	if rw {
		return fmt.Sprintf("\t\t%q: {\n\t\t\tType:             schema.%s,\n\t\t\tOptional:         true,\n\t\t\tDescription:      %q,\n\t\t\tDiffSuppressFunc: AlwaysPresentNotUserProvided,\n\t\t},\n", name, typ, desc)
	}
	return fmt.Sprintf("\t\t%q: {\n\t\t\tType:        schema.%s,\n\t\t\tComputed:    true,\n\t\t\tDescription: %q,\n\t\t},\n", name, typ, desc)
}
