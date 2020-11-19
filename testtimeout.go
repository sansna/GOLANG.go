package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ch := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second * 1)
		ch <- true
	}()

	// which ever comes first gets returned, otherwise blocks here.
	select {
	case <-ctx.Done():
		fmt.Println("ok")
	case val := <-ch:
		fmt.Println(val)
	}
}
