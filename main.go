package main

import (
	"flag"
	"fmt"
	"os"
)

const AppName = "go-template"

func main() {
	fs := flag.NewFlagSet(AppName, flag.ExitOnError)
	initFlags(fs)
	fs.Parse(os.Args[1:])

	fmt.Println("Hello")
	defer fmt.Println("Bye")
}

func initFlags(fs *flag.FlagSet) {

}
