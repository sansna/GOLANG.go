package main
import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i := "-123"
	if string(i[0]) == "-" {
		o, _ := strconv.ParseInt(i, 10, 64)
		fmt.Println(o, reflect.TypeOf(o))
	} else {
		o, _ := strconv.ParseUint(i, 10, 64)
		fmt.Println(o, reflect.TypeOf(o))
	}
}
