package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1.6
	fmt.Println(strconv.FormatFloat(a, 'f', 0, 64))
}
