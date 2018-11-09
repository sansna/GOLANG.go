package main

import "fmt"

func test(rec bool) {
	if (rec) {
		defer func() {
			if r:=recover(); r != nil {
				// fmt.Println("recover condition.")
			} else {
				// fmt.Println("no recover condition.")
			}
		}()
	}
	// panic stops current flow of function
	panic("wer")
	fmt.Println("test: You cannot see this line.")
}

func main() {
	// it recovers from panic
	test(true)
	fmt.Println("main: You can see this line,")
	// no recover from panic, program stopped executing
	test(false)
	fmt.Println("main: But you cannot see this line.")
}
