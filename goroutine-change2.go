package main

import (
	"fmt"
	"time"
)

// test race problem by go run -race x.go
func zzz() {
	// this variable is shared between routines
	var a int
	a = 10
	go func() {
		fmt.Println(a)
		time.Sleep(4 * time.Second)
		fmt.Println(a)
	}()
	time.Sleep(1*time.Second)
	a = 100
	time.Sleep(4*time.Second)
	return
}

func main() {
	zzz()
	fmt.Println("done")
	return
}
