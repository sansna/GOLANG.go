package main
import "strings"
import "fmt"
func main() {
	a := "s,,3,4,"
	strs:=strings.Split(a, ",")
	fmt.Println(len(strs))
	for i:=0; i < len(strs);i ++ {
		fmt.Println(strs[i])
	}
}
