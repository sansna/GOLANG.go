package main

import "fmt"
import "reflect"

func main() {
	var i interface{}
	i = 1
	i = "aslk"
	//var r string
	//r = "hello"
	//i = r
	a, ok:= reflect.ValueOf(i).Interface().(string)
	if (ok == true) {
		fmt.Println(a)
	}
	fmt.Println(reflect.TypeOf(i))
}
