package main

import "fmt"
import "time"

func test() bool {
	return false
}
func main() {
	var a []int
	a = make([]int, 3)
	a[0] = 3
	a[1] = 9
	a[2] = 1
	for index, value := range a {
		fmt.Printf("a%d: %d\n", index, value)
	}
	b := make(map[int64]int64)
	b[1] = 1
	b[0] = 0
	b[1] = 8
	b[4] = 5
	for _, v := range b {
		fmt.Println(v)
	}
	if ok := test(); ok {
		fmt.Println("ok")
	}
	// i is index
	for i := range a {
		fmt.Println(i)
	}
	// v is value
	for _, v := range a {
		fmt.Println(v)
	}

	for i := 0; i != 0; i++ {
		fmt.Println("This will not be printed.")
		fmt.Println(i)
	}

	for true {
		fmt.Println("true")
		time.Sleep(time.Second)
	}
}
