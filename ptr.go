package main
import "fmt"
import "reflect"

type a struct {
	a *int
	b int
}

func main() {
	// p is same as q of type ptr to struct a
	p := &a{
		b:2,
	}
	q := new(a)
	fmt.Println(p.b)
	fmt.Println(p.a)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(q))

	ptrs := make([]*int, 10)
	for i:=0; i< 10; i++ {
		status := 2+i
		ptr := &status
		ptrs[i] = ptr
	}
	fmt.Println(*ptrs[1])
}
