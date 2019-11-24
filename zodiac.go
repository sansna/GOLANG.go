package main

import (
	"fmt"
)

func main() {
	var ZodiacMatchMap = map[int]bool{}
	zpos := []int{1, 3, 5, 9, 11}
	for i := 1; i < 13; i++ {
		for idx, j := range zpos {
			ZodiacMatchMap[i*12+j] = true
			zpos[idx] += 1
			if zpos[idx] > 12 {
				zpos[idx] = 1
			}
		}
	}
	for i := 1; i < 13; i++ {
		for j := 1; j < 13; j++ {
			if ZodiacMatchMap[i*12+j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
