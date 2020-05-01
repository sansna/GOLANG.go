package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// this can read file as binary.
	byt, _ := ioutil.ReadFile("a")
	fmt.Println(byt)
}
