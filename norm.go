package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	n := distuv.Normal{
		Mu:    16,
		Sigma: 1.4,
	}
	fmt.Println(n.CDF(-1))
	fmt.Println(n.Prob(15)/n.Prob(16))
	fmt.Println(n.Prob(14)/n.Prob(16))
	fmt.Println(n.Prob(13)/n.Prob(16))
}
