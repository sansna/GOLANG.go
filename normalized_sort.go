package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Ranker struct {
	Items  []interface{}
	Scores []float64
}

func NewRanker(items []interface{}, scores []float64) *Ranker {
	return &Ranker{
		Items:  items,
		Scores: scores,
	}
}

func NewRankerNormal(items interface{}, scores []float64) *Ranker {
	v := reflect.ValueOf(items)
	is := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		is[i] = v.Index(i).Interface()
	}
	return &Ranker{
		Items:  is,
		Scores: scores,
	}
}

func (r *Ranker) Less(i, j int) bool {
	if r.Scores[i] < r.Scores[j] {
		return true
	}
	return false
}

func (r *Ranker) Len() int {
	return len(r.Scores)
}

func (r *Ranker) Swap(i, j int) {
	r.Scores[i], r.Scores[j] = r.Scores[j], r.Scores[i]
	r.Items[i], r.Items[j] = r.Items[j], r.Items[i]
}

func main() {
	al := []int64{1, 2, 3, 45, 6}
	bl := []float64{5.0, 4.9, 3.8, 1.0, 9.0}

	//is := make([]interface{}, len(bl))
	//for i, v := range al {
	//	fmt.Println(i, v)
	//	is[i] = v
	//}
	//rank := NewRanker(is, bl)

	rank := NewRankerNormal(al, bl)

	fmt.Println(rank)
	sort.Sort(rank)
	fmt.Println(rank)
}
