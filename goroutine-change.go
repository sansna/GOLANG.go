package main

import (
	"fmt"
	"time"
)

// test race problem by go run -race x.go
func zzz() {
	var a int
	a = 10
	go func() {
		fmt.Println("start")
		time.Sleep(4 * time.Second)
		a = 100
		fmt.Println("end")
	}()
	fmt.Println(a)
	time.Sleep(4500*time.Millisecond)
	fmt.Println(a)
	return
}

func main() {
	zzz()
	fmt.Println("done")
	return
}
