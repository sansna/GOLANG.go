package main
import "fmt"
func test() (a,b int) {
	a = 0
	b = 0
	return
}
func main() {
	// outer i, j
	var i,j int
	for i,j = test(); i < 10; i++ {
		fmt.Println(i)
		fmt.Println(j)
	}
	fmt.Println(i)
	// inner i, not affecting outer i
	for i:=0; i< 6; i++ {
		fmt.Println(i)
	}
	// prints outer i
	fmt.Println(i)
}

