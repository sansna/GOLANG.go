package main
import "git.ixiaochuan.cn/service/post/calculator/models/decoder"
import "encoding/json"
import "fmt"
func main() {
	a := &decoder.ActionRedisMQUpdate{
		Pid: 1,
		Age: 2,
		Gender: 1,
		Citycode: "Asdfas",
	}
	b, _ := json.Marshal(a)
	d := string(b)
	fmt.Println(d)
	c := &decoder.ActionRedisMQUpdate{}
	_ = json.Unmarshal([]byte(d), c)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(c.Age)
	fmt.Println(c.Gender)
}
