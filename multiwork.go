package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{}, 3)
	var a, b, c int
	go func() {
		a = 1
		ch <- struct{}{}
	}()
	go func() {
		b = 2
		ch <- struct{}{}
	}()
	go func() {
		c = 3
		ch <- struct{}{}
	}()
	<-ch
	fmt.Println(a, b, c)
	<-ch
	<-ch
	fmt.Println(a, b, c)
}
