package main

import "fmt"
import "math/rand"
import "time"

func genarr(n int) []int {
	rand.Seed(time.Now().UnixNano())
	out := make([]int, n)
	for i := int(0); i < n; i++ {
		out[i] = rand.Intn(10)
	}
	return out
}

func main() {
	a := genarr(10)

	fmt.Println(a)
	quick(a, 0, len(a)-1)

	fmt.Println(a)
}

func quick(arr []int, l, r int) {
	if l < r {
		pos := findPos(arr, l, r)
		quick(arr, l, pos-1)
		quick(arr, pos+1, r)
	}
}

func findPos(arr []int, l, r int) int {
	v := arr[l]
	for l < r {
		for l < r && arr[r] >= v {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= v {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = v
	return l
}
