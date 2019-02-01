package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "今天天气hao\n还行"
	if strings.HasPrefix(a, "今天") {
		fmt.Println("ok")
	} else {
		fmt.Println("err")
	}
	b := strings.Index(a, "\n")
	c := a[b+1:]
	fmt.Println(c)
}
