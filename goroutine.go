package main

import (
	"fmt"
	"time"
)

func zzz() {
	go func() {
		fmt.Println("start")
		time.Sleep(4 * time.Second)
		fmt.Println("end")
	}()
	return
}

func main() {
	zzz()
	fmt.Println("out of zzz now.")
	time.Sleep(time.Second * 5)
	fmt.Println("done")
	return
}
