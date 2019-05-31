package main

import "unicode"
import "fmt"
import "reflect"

func IsAsciiPrintable(s string) bool {
	for _, r := range s {
		// string is composed of runes, and rune is *same* as int32
		//fmt.Println(reflect.TypeOf(r))
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsAsciiPrintable("len([]rune(s)) = %d, len([]byte(s)) = %d\n"))
}
