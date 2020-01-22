package main

import (
	"fmt"
)

func main() {
	map_int := make(map[string]int, 0)
	map_interface := make(map[string]interface{}, 0)
	map_string := make(map[string]string, 0)
	map_float32 := make(map[string]float32, 0)
	map_slice := make(map[string][]int64, 0)
	map_bool := make(map[string]bool, 0)

	fmt.Println(map_int[""],map_bool[""],map_slice[""], map_string[""], map_float32[""], map_interface[""])
}
