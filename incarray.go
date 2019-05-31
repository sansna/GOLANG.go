package main

import "fmt"

func main() {
	var a [2]int64
	a[0] += 1
	a[0] += 1
	a[0] += 1
	a[0] += 1
	a[0] += 1
	a[1] += 1

	// this cannot be done, because the map values are essentially immutable.
	// One have to copy the value out and modify and copy back to have it done.
	//var b map[string][2]int64
	//b = make(map[string][2]int64, 1)
	//b["sdf"] = [2]int64{}
	//b["sdf"][0] += 1
	//b["sdf"][0] += 1
	//b["sdf"][0] += 1
	//b["sdf"][0] += 1
	//b["sdf"][1] += 1

	fmt.Println(a)
	fmt.Println(b)
}
