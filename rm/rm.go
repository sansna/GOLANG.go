package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// only run cmds in shell will have the wildcards in effect (expand to filename list)
	cmd := exec.Command("/usr/local/bin/bash", "-c", "rm -frd *.txt")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("err: %v", err)
	} else {
		fmt.Println("ok")
	}
}
