package main

import (
	"fmt"
	"reflect"
	"math"
)

func main() {
	// on amd64 machine, int32 != (int == int64)
	var a int32
	var b int64
	a=math.MaxInt32
	// overflows to MinInt32
	a++
	// okay
	b=3000000000000000
	// assign a's value to b
	b=int64(a)
	//error overflow:
	//a=3000000000000000
	b=3000000000000000
	// this is not okay, b's value should be checked first.
	a = int32(b)
	fmt.Println(a,b)
	fmt.Println(reflect.TypeOf(a),reflect.TypeOf(b))
}
