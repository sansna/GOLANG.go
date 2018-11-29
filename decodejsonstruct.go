package main

//import "git.ixiaochuan.cn/service/post/calculator/models/decoder"
import "encoding/json"
import "fmt"

type ActionRedisMQUpdate struct {
	Pid            int64  `json:"pid"`
	Tid            int64  `json:"tid,omitempty"`
	Age            int64  //`json:"age"`
	Gender         int64  //`json:"gender"`
	Citycode       string `json:"cc"`
	EmbeddedThings struct {
		Act int64 `json:"act"`
		Adk int64 //`json:"act"`
	} `json:"tag"`
	Intslice []int64 `json:"slice"`
	// the following omitempty tag does not work
	AlwaysInSlice [2]int `json:"always,omitempty"`
	Fixedslice [2]int `json:"fix,omitempty"`
}

func main() {
	//a := &decoder.ActionRedisMQUpdate{
	a := &ActionRedisMQUpdate{
		Pid:      0,
		Age:      2,
		Gender:   1,
		Citycode: "Asdfas",
	}
	a.EmbeddedThings.Act = 32
	a.EmbeddedThings.Adk = 33
	{
		slice := make([]int64, 4)
		slice[0] = 100
		slice[1] = 300
		slice[2] = 302
		a.Intslice = slice
	}
	a.Fixedslice[1] = 10

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
