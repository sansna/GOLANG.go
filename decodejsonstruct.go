package main
//import "git.ixiaochuan.cn/service/post/calculator/models/decoder"
import "encoding/json"
import "fmt"
type ActionRedisMQUpdate struct {
	Pid int64 `json:"pid"`
	Age int64 //`json:"age"`
	Gender int64 //`json:"gender"`
	Citycode string `json:"cc"`
	EmbeddedThings struct {
		Act int64 `json:"act"`
		Adk int64 //`json:"act"`
	} `json:"tag"`
}
func main() {
	//a := &decoder.ActionRedisMQUpdate{
	a := &ActionRedisMQUpdate{
		Pid: 1,
		Age: 2,
		Gender: 1,
		Citycode: "Asdfas",
	}
	a.EmbeddedThings.Act = 32
	a.EmbeddedThings.Adk = 33

	b, _ := json.Marshal(a)
	d := string(b)
	fmt.Println("below decoded message")
	fmt.Println("[]byte message:", b)
	fmt.Println("string message:", d)
	//c := &decoder.ActionRedisMQUpdate{}
	c := &ActionRedisMQUpdate{}
	_ = json.Unmarshal([]byte(d), c)
	fmt.Println("below decoded")
	fmt.Println(c)
	fmt.Println(c.Age)
	fmt.Println(c.Gender)
}
