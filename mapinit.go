package main

import (
	"fmt"
)

func main() {
	var a = map[int64]bool{
		2: true,
	}
	fmt.Println(a[1])
	fmt.Println(len(a))
	var b = map[string]interface{}{
		"1": 1,
	}
	fmt.Println(b)
	fmt.Println(b["1"])
	fmt.Println(b["2"])
	// This works for non-existence cases
	fmt.Println(b["1"] == 1)
}
