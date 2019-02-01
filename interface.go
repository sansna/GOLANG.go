package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	b := make([]interface{}, 0)
	b = append(b, "sdf")
	b = append(b, 2)
	fmt.Println(a, b)
	a = b
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b[0]), reflect.TypeOf(b[1]))
	fmt.Println(a)
	c := make([]interface{}, 0)
	a = c
	fmt.Println(a, len(a.([]interface{})))
	// c is nil, error
	//fmt.Println(a, len(a.([]interface{})), a.([]interface{})[0])
}
