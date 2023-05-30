package main

import (
	"flag"
	"fmt"
	"os"
)

const AppName = "go-template"

// build info
var (
	version = "development"
	commit  = "N.A."
	date    = "N.A."
)

// flags
var (
	showVersion   *bool
	showBuildInfo *bool
)

func main() {
	fs := flag.NewFlagSet(AppName, flag.ExitOnError)
	initFlags(fs)
	fs.Parse(os.Args[1:])

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if *showBuildInfo {
		fmt.Printf("Version:%s, GitCommit:%s, BuildDate:%s\n", version, commit, date)
		os.Exit(0)
	}
	fmt.Println("Hello")
	defer fmt.Println("Bye")
}

func initFlags(fs *flag.FlagSet) {
	showVersion = fs.Bool("v", false, "Print version and exit")
	showBuildInfo = fs.Bool("V", false, "Print build information and exit")
}
