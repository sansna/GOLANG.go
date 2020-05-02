package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// this should always exists when using flag
	flag.Parse()
	// check args count
	if flag.NArg() == 0 {
		return
	}
	// positional args, all parameters
	midStr := flag.Args()[0]
	fmt.Println(midStr, flag.Args())

	// positional args, as commandline
	fmt.Println(os.Args, os.Args[1])
}
