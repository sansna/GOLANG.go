package main

import "math/rand"
import "time"
import "fmt"

var (
	Count = 0
)

func genarr(n int) []int {
	rand.Seed(time.Now().UnixNano())
	out := make([]int, n)
	for i := int(0); i < n; i++ {
		out[i] = rand.Intn(200000000)
	}
	return out
}

func main() {
	n := 100000000
	dst := make([]int, n)
	sel2 := make([]int, n)
	arr := genarr(n)
	copy(dst, arr)
	copy(sel2, arr)

	length := len(arr)
	max := length / 2
	min := (length - 1) / 2
	posmap := map[int]bool{min: false, max: false}

	//fmt.Println(arr, posmap)
	st1 := time.Now().UnixNano()
	qsort(arr, 0, len(arr)-1, posmap)
	et1 := time.Now().UnixNano()
	//fmt.Println(arr, posmap, Count)
	cnt1 := Count

	//Count = 0

	////fmt.Println(dst, nil)
	//st2 := time.Now().UnixNano()
	//qsort(dst, 0, len(dst)-1, map[int]bool{})
	//et2 := time.Now().UnixNano()
	//cnt2 := Count
	////fmt.Println(dst, nil, Count)

	fmt.Println("优化快排", arr[min], arr[max], cnt1, et1-st1)
	//fmt.Println("快排", dst[min], dst[max], cnt2, et2-st2)
	//fmt.Println("优化时间：", (et2-st2)-(et1-st1))

	//st3 := time.Now().UnixNano()
	//valn := Select(sel2, min+1)
	//valx := Select(sel2, max+1)
	//et3 := time.Now().UnixNano()
	//fmt.Println("sele:", valn, valx, et3-st3)
}

func qsort(arr []int, l, r int, posmap map[int]bool) {
	if l < r {
		if len(posmap) > 0 {
			hit := false
			for i := l; i < r; i++ {
				if v, ok := posmap[i]; ok && !v {
					hit = true
					break
				}
			}
			if !hit {
				return
			}
		}

		Count += 1
		pos := findPos(arr, l, r, posmap)
		if len(posmap) > 0 {
			hit := true
			for _, v := range posmap {
				if !v {
					hit = false
					break
				}
			}
			if hit {
				return
			}
		}
		qsort(arr, l, pos-1, posmap)
		qsort(arr, pos+1, r, posmap)
	}
}

func findPos(arr []int, l, r int, posmap map[int]bool) int {
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

	if len(posmap) > 0 {
		if _, ok := posmap[l]; ok {
			posmap[l] = true
		}
	}
	return l
}

// educational use, not practical in production
func Select(arr []int, k int) int {
	if len(arr) <= 1 {
		return arr[0]
	}
	t := arr[0]
	var s1, s2, s3 []int
	for _, v := range arr {
		if v < t {
			s1 = append(s1, v)
		} else if v > t {
			s3 = append(s3, v)
		} else {
			s2 = append(s2, v)
		}
	}
	if len(s1) >= k {
		return Select(s1, k)
	} else if len(s1)+len(s2) >= k {
		return t
	} else {
		return Select(s3, k-len(s1)-len(s2))
	}
}
