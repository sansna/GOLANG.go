package main

import (
	"fmt"
	"sort"
)

type Aa struct {
	length int
	a      []int64
}

func (aa Aa) Less(i, j int) bool {
	return aa.a[i] > aa.a[j]
}

func (aa Aa) Len() int {
	return int(aa.length)
}

func (aa Aa) Swap(i, j int) {
	aa.a[i], aa.a[j] = aa.a[j], aa.a[i]
}

func main() {
	z := []int64{3, 2, 5, 6, 1, 9}
	zz := Aa{
		length: len(z),
		a:      z,
	}
	// default sort to ascending list, however...see Less.
	sort.Sort(zz)
	fmt.Println(zz)
	fmt.Println(zz.Len())
}
