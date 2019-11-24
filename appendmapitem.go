package main

import (
	"fmt"
)

var (
	AA = map[int64]string{
		1: "zz",
		2: "ik",
	}
)

func main() {
	strs := []string{}
	iis := []int64{1, 2, 3, 4}
	for _, i := range iis {
		// Use this to skip ""
		//if AA[i] == "" {
		//	continue
		//}
		// non-exist map is mapped to ""
		strs = append(strs, AA[i])
	}
	fmt.Println(strs, len(strs))
}
