package main
import "fmt"
import "errors"

func sss() (mids []int64, s string, err error) {
	//mids = []int64{}
	//s = string{}
	err = errors.New("wow")
	return
}

func main() {
	a, b, er := sss()
	fmt.Println(a,b,er)
}
