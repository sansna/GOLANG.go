package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// array is actually a ptr, so the following is okay
	m := map[int64][]int64{}
	m[0] = []int64{1, 2}
	fmt.Println(m, m[0], unsafe.Sizeof(m), unsafe.Sizeof(m[0]))
	m[0] = append(m[0], 3)
	m[0] = append(m[0], 4)
	m[0][1] = 100
	fmt.Println(m, m[0], unsafe.Sizeof(m), unsafe.Sizeof(m[0]))

	//
	type AA struct {
		A int
		B int
		C int
		D int
	}
	n := map[int64]AA{}
	n[0] = AA{
		B: 3,
	}
	fmt.Println(n[0], unsafe.Sizeof(n), unsafe.Sizeof(n[0]))
	// this must fail, AA is as a compact thing to map's value, cannot be modified.
	//n[0].A = 10
	n[0] = AA{
		A: n[0].A,
		B: n[0].B,
		C: 10,
	}
	fmt.Println(n[0], unsafe.Sizeof(n), unsafe.Sizeof(n[0]))

	// Only when array of ptrs, the values are modifiable. []*AA
	// but array of structs, its not modifiable. []AA
	l := make(map[int64][]*AA)
	l[10] = append(l[10], &AA{
		A: 1,
		B: 10,
	})
	l[10] = append(l[10], &AA{
		A: 2,
		B: 20,
	})
	fmt.Println(l[10])
	for _, v := range l {
		for _, item := range v {
			if item.A == 2 {
				fmt.Println("into")
				item.C = 1827
				fmt.Println("into", l[10])
			}
		}
	}
	for _, v := range l[10] {
		fmt.Println(v)
	}
}
