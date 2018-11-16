package main
import "fmt"
func test() (a,b int) {
	a = 0
	b = 0
	return
}
func main() {
	var i,j int
	for i,j = test(); i < 10; i++ {
		fmt.Println(i)
		fmt.Println(j)
	}
	fmt.Println(i+1)
}

