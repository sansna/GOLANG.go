// Testing this code by: go run main.go
package main

import "fmt"
import (
	"math"
)

func main() {
	var a int = 1
	var b, c int = 2, 3
	var d float64
	var arr [10]int             // all to 0
	var arr2 = [10]int{1, 2, 3} // others to 0
	fmt.Println("hello world!")
	a = b + c
	// type cast, losing precision
	d = (float64)(float32(a) / float32(c))
	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(arr[0], arr2[0])

	/*
	 * declaration can be anywhere
	 * multi-array declare & init
	 */
	var marr = [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 1, 0},
	}
	fmt.Println(marr[2][1])

	/*
	 * ptr usage
	 * &a, *a like c.
	 */
	var ptr *int
	ptr = &a
	fmt.Println("location at a: %x", ptr)

	// define struct People and usage of People
	type People struct {
		name  string
		phone string
	}
	var people People
	var people2 People
	var ppeople *People = &people2
	people.name = "tim"
	people.phone = "156"
	ppeople.name = "tam"
	ppeople.phone = "123"
	fmt.Printf("name1 is %s, name2 is %s\n", people.name, ppeople.name)

	// slice & sub-slice
	var slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("%v\n", slice[2:4]) //sub-slice start from 2 and ends at 3. 4 is excluded
	slice2 := make([]int, 2, 11)   //slice2 which has 2 member. Max 11 mem capable.
	slice2[1] = 3
	slice2 = append(slice2, 4)
	fmt.Printf("len%d,cap%d,slice=%v\n", len(slice2), cap(slice2), slice2)
	copy(slice2, slice) //can only copy same number of len(slice2) from slice
	fmt.Printf("len%d,cap%d,slice=%v\n", len(slice2), cap(slice2), slice2)

	// map usages
	//  declare a variable which is nil
	var people_list map[string]string
	//  make the variable to actually be a map
	people_list = make(map[string]string)
	//  insert values
	people_list["0"] = "zhang"
	people_list["1"] = "li"
	people_list["2"] = "liu"
	people_list["3"] = "wang"
	// randomly iterate through people_list
	// both following are okay.
	//	for i := range people_list {
	//		fmt.Println(people_list[i])
	//	}
	for i, j := range people_list {
		fmt.Println(i, j)
	}

	// recursion function call, uncomment it to run infini loop
	//recursion()

	// intfc
	circle := Circle{radix: 3}
	rect := Rect{x: 3, y: 4}
	fmt.Println(circle.area())
	fmt.Println(rect.area())
}

func recursion() {
	recursion()
}

type Circle struct {
	radix float64
}

type Rect struct {
	x float64
	y float64
}

type Shape interface {
	area() float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radix * circle.radix
}

func (rect Rect) area() float64 {
	return rect.x * rect.y
}
