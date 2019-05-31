package main

import "fmt"

func main() {
	long := make(map[string]interface{}, 100)
	fmt.Println("length of long is", len(long))
	m := make(map[string]interface{})
	m["life"] = 1
	m["is"] = 2
	ebd := make(map[string]string)
	m["good"] = ebd
	ebd["inner1"] = "golang"
	ebd["inner2"] = "again"
	fmt.Println(m)

	value, exists := m["love"]
	fmt.Println(value, exists)
	value, exists = m["life"]
	fmt.Println(value, exists)

	mm := map[string]interface{}{}
	mm["li"] = 13
	fmt.Println(mm)
	if _, ok := mm["sd"]; ok {
		fmt.Println("oksd")
	}
	if _, ok := mm["li"]; ok {
		fmt.Println("okli")
	}

	zh := "sldf"
	mmm := map[string]interface{}{
		"Lif": 1,
		"zh":  zh,
		"mm":  mm,
	}
	fmt.Println("mmm", mmm)
	fmt.Println(len(mmm))
	delete(mmm, "Lif")
	fmt.Println("mmm", mmm)
	fmt.Println(len(mmm))

	for n, v := range mmm {
		fmt.Println("n:", n, "v:", v)
	}

	// zero v will also be iterated.
	nn := map[int64]int64{
		1: 0,
	}
	nn[123] = 0
	for n, v := range nn {
		fmt.Println(n, v)
	}

	// map of array
	array := []int64{3, 5, 7}
	moa := map[int64]*[]int64{
		1: &[]int64{1, 2, 3, 5},
		3: &[]int64{2, 3, 5},
		5: &array,
	}
	array = append(array, 1100)
	*moa[1] = append(*moa[1], 102)
	(*moa[3])[1] = 0
	fmt.Println("moa:", *moa[1], *moa[3], *moa[5])


	var newmap map[string]string
	newmap = nil
	for k, v := range newmap {
		fmt.Println("info:",k,v)
	}

	if _, ok:=newmap["asdf"];!ok {
		fmt.Println("asdf not found")
	}

	// create instance of map
	fmt.Println(map[int64][2]int64{
		1: [2]int64{2,3},
	})
}
