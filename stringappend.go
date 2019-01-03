package main
import "fmt"
func main() {
	var a string
	a += "als"
	a += "als"
	var b string
	fmt.Println(a)
	fmt.Println(b)
	c := a+b+a
	fmt.Println(c)
}
