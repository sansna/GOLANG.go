package main
import (
	"fmt"
	"reflect"
)

type Zaa struct {
	a int64
	b int64
}

func main() {
	z := Zaa{
		a: 1,
		b: 2,
	}
	fmt.Println(reflect.TypeOf(z).String()=="main.Zaa")

	b := Zaa{}
	if b == (Zaa{}) {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
