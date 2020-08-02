package main

import (
	"fmt"
)

func main() {
	var a interface{}
	a = int64(1)
	if _, ok := a.(string); ok {
		fmt.Println(a, a.(int64))
	}
	if _, ok := a.(float64); ok {
		fmt.Println(a, a.(int64))
	}
	if _, ok := a.(int); ok {
		fmt.Println(a, a.(int64))
	}
	// only this works
	if _, ok := a.(int64); ok {
		fmt.Println(a, a.(int64))
	}

	//
	m := map[string]interface{}{
		"asdf": int64(1),
	}
	// this does exactly what we need.
	if _, ok := m["asdf"].(int64); ok {
		fmt.Println(m)
	}

	// this leads to panic
	//aid := m["aid"].(int)
	//fmt.Println(aid)
}
