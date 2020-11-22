package main

import "fmt"
import "git.ixiaochuan.cn/pp_server/live/util"

type SA struct{}
type SB struct{}

func (SB) GetConst() int {
	return 100
}

func (SA) GetConst() int {
	return 200
}

func (s SA) GetConsta() string {
	return fmt.Sprintf("%p", &s)
}

func (s SB) GetConsta() string {
	return fmt.Sprintf("%p", &s)
}

func gb() int {
	a := SB{}
	return a.GetConst()
}

func ga() int {
	return SA{}.GetConst()
}

func main() {
	fmt.Println(gb())
	fmt.Println(ga())
	// the following print same, because all struct{} instance share the same obj.
	// but each one's GetConsta() method differs.
	c := util.SA{}
	fmt.Println(c.GetConsta())
	fmt.Println(SA{}.GetConsta())
	fmt.Println(SB{}.GetConsta())
}
