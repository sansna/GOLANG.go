package main
import "fmt"
func sss() int {
	return 10
}
func main() {
	var i int
	for {
		// this does *not* affect outer i value
		//i := sss()
		// this affect outer i value
		i = sss()
		if i == 10 {
			break
		}
	}
	fmt.Println(i)
}
