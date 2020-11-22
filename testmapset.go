package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// much less in size storage ram
	maps := make(map[int64]struct{})
	// much worse case when implementing set in go
	mapb := make(map[int64]bool)
	for i := 0; i < 10; i++ {
		maps[int64(i)] = struct{}{}
		mapb[int64(i)] = false
	}
	fmt.Println("s:", unsafe.Sizeof(maps))
	fmt.Println("b:", unsafe.Sizeof(mapb))

	for k := range maps {
		fmt.Println(k, maps[k])
		fmt.Println(k, mapb[k])
	}

}
