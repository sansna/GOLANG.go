package main
import "fmt"
func arrayappend(result *[]*int) {
	a := 1
	b := 2
	*result = append(*result, &a)
	*result = append(*result, &b)
}

func main() {
	result := make([]*int,0,10)
	arrayappend(&result)
	fmt.Println(*result[0],*result[1])
}
