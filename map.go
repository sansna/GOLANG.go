package main
import "fmt"
func main() {
	m := make(map[string]interface{})
	m["life"]=1
	m["is"]=2
	ebd := make(map[string]string)
	m["good"]=ebd
	ebd["inner1"]="golang"
	ebd["inner2"]="again"
	fmt.Println(m)

	value, exists := m["love"]
	fmt.Println(value,exists)
	value, exists = m["life"]
	fmt.Println(value,exists)
}
