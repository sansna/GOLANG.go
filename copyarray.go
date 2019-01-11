package main
import "fmt"
func main() {
	a := make([]int64, 0, 2)
	a = append(a, 1)
	a = append(a, 2)

	b := a
	b[1] = 3
	// b is by reference to a!
	fmt.Println(a)

	a[1] = 2
	c := make([]int64, 0, 2)
	copy(c[:], a[:])
	// c has no len, thus nothing copies
	fmt.Println(c)
	
	a[1] = 2
	// true way of doing copy by value
	d := make([]int64, 2)
	copy(d[:], a[:])
	fmt.Println(d)
}
