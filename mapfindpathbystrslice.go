package main

import (
	"fmt"
)

func main() {
	m := map[string]interface{}{}
	//var t_m map[string]interface{}
	ss := []string{"asdf", "Zx", "s"}
	iter_m := m
	for _, s := range ss {
		if _, ok := iter_m[s]; !ok {
			iter_m[s] = make(map[string]interface{})
		}
		iter_m = iter_m[s].(map[string]interface{})
	}
	another_m := map[string]interface{}{
		"1":   123,
		"zxc": "asdf",
	}
	for k, v := range another_m {
		iter_m[k] = v
	}
	fmt.Println(m)
	fmt.Println(iter_m)

	return
}
