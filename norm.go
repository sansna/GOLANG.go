package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	n := distuv.Normal{
		Mu:    0,
		Sigma: 1,
	}
	fmt.Println(n.CDF(0))
}
