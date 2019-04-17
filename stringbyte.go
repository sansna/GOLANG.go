package main

import (
	"fmt"
)

func main() {
	s := "asdf"
	b := []byte(s)
	fmt.Println(len(b), len(s))
}
