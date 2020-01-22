package main

import (
	"fmt"
)

func main() {
	a := "1000123812478"
	fmt.Println(a[len(a)-4 : len(a)])
}
