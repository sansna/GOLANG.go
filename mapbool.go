package main

import (
	"fmt"
)

func main() {
	mapbool := make(map[int64]bool, 0)

	// yeah. this makes differences.
	mapbool[1] = false
	mapbool[2] = true
	fmt.Println(mapbool)
	for k, v := range mapbool {
		fmt.Println(k, v)
	}


	mapbool2 := make(map[int64]bool, 3)
	fmt.Println(mapbool2)
}
