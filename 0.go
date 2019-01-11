package main
import "fmt"
func main() {
	a := uint32(0)
	a = ^a
	fmt.Printf("%032b", a)
}
