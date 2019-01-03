package main
import "fmt"
import "time"
type bbb struct {
	a int
	b int
}
type aaa struct {
	bbb
	time.Duration
	c int
}
func main() {
	e := bbb{
		a:1,
		b:2,
	}
	d := aaa{
		bbb: e,
		time.Duration: time.Second,
		c: 2,
	}
	fmt.Println(d)
}
