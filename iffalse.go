package main

import (
	"time"
	"fmt"
)

func main() {
	go func() {
		if false {
			fmt.Print("ok")
		}
	}()
	time.Sleep(time.Second)
}
