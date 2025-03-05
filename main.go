package main

import (
	"flag"
	"os"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}
}
