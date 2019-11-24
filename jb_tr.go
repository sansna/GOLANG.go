package main

import (
	"fmt"

	"github.com/wangbin/jiebago/analyse"
)

func main() {
	var tr analyse.TextRanker
	tr.LoadDictionary("dict.txt")
	ret := tr.TextRank("我是一个我说没有地球明天的人", 10)
	for _, r := range ret {
		fmt.Println(r.Text())
	}
}
