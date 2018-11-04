package main

import "fmt"
import "reflect"

func main() {
	type Base struct {
		baseint int
	}
	// annonymous member of struct Base, its members are automatically
	// added to struct Elt.
	type Elt struct {
		*Base
		id, phone int
		name string
	}
	// example of init annonymous member during allocation
	a := Elt{
		Base: &Base{
			baseint: 10,
		},
		id: 1,
		phone: 12,
		name: "Asdf",
	}
	// ptr low efficiency, on heap
	b := &a
	// value on stack, more effi, more ram
	c := a
	// c.Base = &Base{}
	b.id = 2
	// annonymous type direct call member
	a.baseint = 11

	fmt.Println(a.id)
	fmt.Println(c.id)
	// ptr/value both use dot as member dereference
	fmt.Println(reflect.TypeOf(b.id))
	fmt.Println(&b, &(b.id))
	
	// because Base is of ptr type,
	// change of a/b's Base elts also changes c's Base elts
	fmt.Println(c.baseint)
}
