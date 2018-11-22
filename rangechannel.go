package main
import "fmt"
import "time"
func main() {
	c := make(chan int,100)
	go func() {
		for a := range c {
			fmt.Println(a)
		}
	}()
	for _= range time.Tick(time.Second * 2) {
		c<- 3
	}
}
