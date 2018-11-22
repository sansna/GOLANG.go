package main
import "fmt"
func main() {
	var b string
	a := "helloworld"
	b = a
	a = "Asdf"
	fmt.Println(&a, &b,a,b)
}
