package main
import "fmt"
type a struct {
	a int64
}
func main() {
	// allocate slot for slices of ptrs
	b := make([]*a,10)
	// allocate ptr addr for a ptr
	b[1] = new(a)
	b[1].a =10
	fmt.Println(b[1].a)
}
