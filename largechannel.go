package main

import "fmt"
import "time"
import "runtime"

func main() {
	runtime.GOMAXPROCS(3)
	c := make(chan int, 3000000000)
	go func() {
		for range time.Tick(time.Second) {
			// NOTE: this memory may be collected by go's GC
			fmt.Printf("len of chan: %d, size of mem: %dMB\n", len(c), len(c)*8/1024/1024)
		}
	}()
	go func() {
		time.Sleep(time.Second * 100)
		a := <-c
		fmt.Println("a ok:", a)
		return
	}()
	for {
		c <- 1
	}
	fmt.Println("ok")
	return
}
