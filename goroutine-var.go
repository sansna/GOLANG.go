package main

import (
	"fmt"
	"time"
)

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
}
