package main

import (
	"fmt"
	"github.com/go-audio/wav"
	"os"
)

func main() {
	f, _ := os.Open("/Users/sansna/Desktop/a.wav")
	defer f.Close()
	w := wav.NewDecoder(f)
	fmt.Println(w.Duration())
}
