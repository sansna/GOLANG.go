package main
import "fmt"
var	m = map[int64]bool{
		1: false,
		2: true,
		3: false,
		4: true,
	}
func main() {
	a := int64(2)
	if m[a] {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
