package main
/* 
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -la
#include "a.h"
*/
import "C"
//import "fmt"

func main() {
	C.Lib()
	return
}
