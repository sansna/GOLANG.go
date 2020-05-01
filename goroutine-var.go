package main

import (
	"fmt"
	"runtime"
	"time"
)

func testlife() {
	// lifetime
	z := 1
	// z'val is preserved after testlifes's exit, even after GC-ed
	go func() {
		for true {
			time.Sleep(time.Millisecond * 20)
			fmt.Println(z)
		}
	}()
	time.Sleep(time.Second * 3)
	// run gc
	runtime.GC()
	return
}

func main() {
	// This does not work.
	//go func() {
	//	a := 10
	//}()
	//fmt.Println(a)

	// This defines b outside, can be used inside, and also changes outside!
	b := 10
	go func() {
		b = 100
	}()
	fmt.Println("first:", b)
	time.Sleep(time.Second * 2)
	fmt.Println("second:", b)

	fmt.Println("go into testlife")
	testlife()
	fmt.Println("out of testlife")
	runtime.GC()

	time.Sleep(time.Second * 10)
	return
}
