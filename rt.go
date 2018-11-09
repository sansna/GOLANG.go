package main

import "fmt"
import "time"
import "runtime"
import "runtime/debug"

func main() {
	fmt.Println(runtime.GOMAXPROCS(2))
	fmt.Println(runtime.GOMAXPROCS(0))
	chblk := make(chan int)
	//ch100 := make(chan int, 100)
	debug.SetMaxThreads(12)
	runtime.GOMAXPROCS(10)
	for {
		go func() {
			chblk<- 3
			time.Sleep(time.Second*4)
		}()
		if runtime.NumGoroutine() > 10 {
			break
		}
	}
	v, _:= <-chblk
	//ch100<- 4
	//v100, _ := <-ch100
	fmt.Println(v,runtime.NumGoroutine())
	fmt.Println(runtime.GOROOT())
	time.Sleep(time.Second*4)
}
