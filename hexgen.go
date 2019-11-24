package main

import (
	"fmt"
	"strconv"
)

func main() {
	s, _ := strconv.ParseInt("fe20", 16, 64)
	e, _ := strconv.ParseInt("fe2f", 16, 64)
	for i := s; i <= e; i++ {
		fmt.Print("[\\u" + strconv.FormatInt(i, 16) + "]*")
	}
}
