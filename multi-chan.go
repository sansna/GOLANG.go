package main
import "fmt"
var c = make(chan int)

func f(n int) {
	for val := range c {    
		fmt.Printf("routine %d : %v\n", n, val)
	}

}   

func main() {
	go f(1)
	go f(2)
	var i int
	i = 0
	for {
		i ++
		c <- i
	}

	close(c) 
}
