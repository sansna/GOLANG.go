package main
import (
	"fmt"
	"runtime"
)
func sss() {
	fpcs := make([]uintptr, 1)
	n := runtime.Callers(1, fpcs)
	if n == 0 {
		fmt.Println("n/a")
	} else {
		fun := runtime.FuncForPC(fpcs[0]-1)
		fmt.Println(fun.Name())
	}
}

func ss1() {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Println(file, no)
	}
}

func main() {
	sss()
	fmt.Println("another way")
	ss1()
	return
}
