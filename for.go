package main
import "fmt"
func main() {
	var a []int
	a = make([]int,3)
	a[0] = 3
	a[1] = 9
	a[2] = 1
	for index,value := range a {
		fmt.Println("a%d: %d", index, value)
	}
	b:=make(map[int64]int64)
	b[1]=1
	b[0]=0
	b[1]=8
	b[4]=5
	for _, v := range b {
		fmt.Println(v)
	}
}
