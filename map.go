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
}
