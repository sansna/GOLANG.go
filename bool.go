package main

import "fmt"

func main() {
	var a bool
	a = true
	a = a | false
	if a {
		fmt.Println("ok")
	}
}
