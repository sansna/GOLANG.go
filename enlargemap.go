package main

import (
	"fmt"
)

func main() {
	t := make(map[string]bool, 0)
	t["!23"] = true
	t["xx"] = true
	t["x"] = false
	fmt.Println(len(t))
	return
}
