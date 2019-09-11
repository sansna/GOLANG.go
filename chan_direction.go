package main

import (
	"fmt"
	"time"
)

var a chan string

func getOutput() <-chan string {
	return a
}

func getInput() chan<- string {
	return a
}

func main() {
	a = make(chan string)
	go input()
	go output()
	time.Sleep(time.Minute * 1)
	return
}

func input() {
	in := getInput()
	for range time.Tick(time.Second * 1) {
		in <- "asdfasdfk"
	}
	return
}

func output() {
	out := getOutput()
	for {
		select {
		case v := <-out:
			fmt.Println(v)
		}
	}
	return
}
