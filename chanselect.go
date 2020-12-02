package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(1)

	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)

	ch1 <- 1
	ch2 <- "hello"

	select {
	// randomly select one
	case <-ch1:
		fmt.Println("ch1")
		return
	case <-ch2:
		fmt.Println("ch2")
		return
	}
}
