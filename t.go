package main

import "time"
import "fmt"

func main() {
	fmt.Printf("%d\n", time.Now().Month())
	fmt.Printf("%d\n", time.Now().Unix()%3600/300*300)
}
