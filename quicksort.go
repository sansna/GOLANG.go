package main

import "fmt"

func main() {
	a := []int{1,2,4,5,6,7}

	quick(a, 0, len(a)-1)

	fmt.Println(a)
}

func quick(arr []int, l, r int) {
	t := arr[l]
	ptr := l

	i, j := l, r

	for i < j {
		for j >= ptr && arr[j] > arr[ptr] {
			j--
		}

		if j >= ptr {
			arr[ptr] = arr[j]
			ptr = j
		}

		for i <= ptr && arr[i] <= arr[ptr] {
			i++
		}

		if i <= ptr {
			arr[ptr] = arr[i]
			ptr = i
		}

		arr[ptr] = t

		if ptr-l > 1 {
			quick(arr, l, ptr-1)
		}
		if r-ptr > 1 {
			quick(arr, ptr+1, r)
		}
	}
	fmt.Println(l, r)
}
