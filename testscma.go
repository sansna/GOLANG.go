package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool, 4)
	go func() {
		<-ctx.Done()
		fmt.Println("ok")
		done <- true
	}()
	go func() {
		<-ctx.Done()
		fmt.Println("ok")
		done <- true
	}()
	go func() {
		<-ctx.Done()
		fmt.Println("ok")
		done <- true
	}()

	cancel()

	<-done
	<-done
	<-done
	fmt.Println("zz")
}
