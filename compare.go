package main
import (
	"fmt"
)
func main() {
	a := "sd"
	var b interface{}
	c := []byte("sd")
	b = c
	fmt.Println(c,b)
	if (b) == a {
		fmt.Printf("%v", b)
	}
}
