package main

import "fmt"

type People struct {
}
type Teacher struct {
	People
}

func (p People) ShowA() {
	fmt.Println("asdf")
	p.ShowB()
}

func (t Teacher) ShowB() {
	fmt.Println("tshow b")
}

func (People) ShowB() {
	fmt.Println("showb")
}

func main() {
	t := &Teacher{}
	t.ShowA()
}
