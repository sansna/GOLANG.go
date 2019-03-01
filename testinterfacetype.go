package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	b := "asdf"
	a = b
	fmt.Println(reflect.TypeOf(a))

	if reflect.TypeOf(a).Kind() == reflect.String {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
	//if len(a.([]interface{})) != 2 {
	//	fmt.Println("len != 2")
	//}

	c := []string{"!2", "Asdf"}
	a = c
	fmt.Println(reflect.TypeOf(a))

	if reflect.TypeOf(a).Kind() == reflect.String {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
	if len(a.([]interface{})) != 2 {
		fmt.Println("len != 2")
	}
}
