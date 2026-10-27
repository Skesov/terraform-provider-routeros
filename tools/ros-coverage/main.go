// ros-coverage compares the RouterOS menu tree of a live router with the
// provider schema.
//
//	go run ./tools/ros-coverage dump -url https://router -o tools/ros-coverage/schema-7.25-board.json
//	go run ./tools/ros-coverage compare tools/ros-coverage/schema-*.json
//	go run ./tools/ros-coverage gen -resource routeros_ip_ssh_server tools/ros-coverage/schema-*.json
//	go run ./tools/ros-coverage gen -path /ip/proxy -name routeros_ip_proxy tools/ros-coverage/schema-*.json
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	switch os.Args[1] {
	case "dump":
		fs := flag.NewFlagSet("dump", flag.ExitOnError)
		url := fs.String("url", os.Getenv("ROS_HOSTURL"), "router URL, https://host")
		out := fs.String("o", "", "output file (default schema-<version>-<board>.json)")
		_ = fs.Parse(os.Args[2:])
		if err := runDump(*url, *out); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "compare":
		fs := flag.NewFlagSet("compare", flag.ExitOnError)
		filter := fs.String("f", "", "only paths containing this string")
		used := fs.Bool("configured", false, "only menus with configuration on the router")
		details := fs.Bool("details", true, "list properties per menu")
		_ = fs.Parse(os.Args[2:])
		if err := runCompare(fs.Args(), *filter, *used, *details); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "gen":
		fs := flag.NewFlagSet("gen", flag.ExitOnError)
		res := fs.String("resource", "", "existing resource: print entries for its missing properties")
		path := fs.String("path", "", "menu without a resource: write a draft resource")
		name := fs.String("name", "", "resource name for -path, routeros_...")
		repo := fs.String("repo", ".", "provider repository root for -path")
		_ = fs.Parse(os.Args[2:])
		var err error
		if *res != "" {
			err = runGen(fs.Args(), *res)
		} else {
			err = runGenResource(fs.Args(), *repo, *path, *name)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		usage()
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: ros-coverage dump -url URL [-o FILE] | compare [-f PATH] [-configured] FILE... | gen (-resource NAME | -path PATH -name NAME) FILE...")
	os.Exit(2)
}
