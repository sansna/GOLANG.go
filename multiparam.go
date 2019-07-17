package main

import "fmt"
import "reflect"

func mp(ss ...string) string {
	var out string
	for _, s := range ss {
		out += s
	}
	fmt.Println(len(ss), reflect.TypeOf(ss))
	return out
}

func main() {
	fmt.Println(mp("sdf"))
	fmt.Println(mp("sdf", "zxcv"))
	fmt.Println(mp("sdf", "ZXcv", "你好"))
	fmt.Println(mp("sdf", "13", "xx"))
	//fmt.Println(mp("sdf", 13, 133))
}
