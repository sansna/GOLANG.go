package main
/* 
#cgo CFLAGS: -I/home/trii/so
#cgo LDFLAGS: -L/home/trii/so -la
#include "a.h"
*/
import "C"
//import "fmt"

func main() {
	C.Lib()
	return
}
