package main

import "fmt"

func block_chan() {
	c:=make(chan int)
	defer close(c)
	// channel is blocked both w/r side,
	// so a goroutine is needed.
	go func() {
		c<-2
	}()
	a := <-c
	fmt.Println("block_chan: ", a)
}

func nonblock_chan() {
	c:=make(chan int, 1)
	defer close(c)
	// channel is nonblock,
	// so a goroutine is not coearce.
	// but if buffer gets full, it is still blocked.
	c<-2
	a := <-c
	fmt.Println("nonblock_chan: ", a)
}

func main() {
	block_chan()
	nonblock_chan()
}
