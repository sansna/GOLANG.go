package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	v int
}

func main() {
	d := &Data{
		v: 10,
	}
	z := d.ret()
	// prints *Data other than interface{}
	fmt.Println(reflect.TypeOf(z))
}

func (d *Data) ret() interface{} {
	return d
}
