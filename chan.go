package main

import (
	"fmt"
	"reflect"
)

func block_chan() {
	var zzz chan interface{}
	c := make(chan *int)
	defer close(c)
	// channel is blocked both w/r side,
	// so a goroutine is needed.
	z := 2
	go func() {
		c <- &z
	}()
	a := <-c
	fmt.Println("block_chan: ", *a)
	fmt.Println("block_chan: ", reflect.TypeOf(a))
	fmt.Println("block_chan: ", reflect.TypeOf(a).String() == "*int")

	// chan detect
	if zzz == nil {
		fmt.Println("ok")
	} else {
		fmt.Println("sd")
	}
}

func nonblock_chan() {
	c := make(chan int, 1)
	defer close(c)
	// channel is nonblock,
	// so a goroutine is not coearce.
	// but if buffer gets full, it is still blocked.
	c <- 2
	a := <-c
	fmt.Println("nonblock_chan: ", a)
}

// channel can return as directional chan, but cannot reversely.
func return_direct_chan() <-chan string {
	var a chan string
	a = make(chan string, 1000)
	return a
}

func main() {
	block_chan()
	nonblock_chan()
	d := return_direct_chan()
	if d != nil {
		fmt.Println("ok")
	}
}
