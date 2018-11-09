package main
import "time"
import "fmt"
func main() {
	a := time.NewTimer(time.Second*3)
	fmt.Println(time.Now().Unix())
	b := <-a.C
	fmt.Println(b.Unix())
	fmt.Println(time.Now().Unix())
}
