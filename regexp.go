package main

import "regexp"
import "fmt"

func main() {
	reg1 := regexp.MustCompile("(\\d{5,})|([a-zA-Z]{1,}[0-9_-]{1,})")
	a := "mag-12"
	if reg1.Match([]byte(a)) {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
