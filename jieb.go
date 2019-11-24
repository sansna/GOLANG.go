package main

import (
	"bufio"
	"fmt"
	"github.com/wangbin/jiebago"
	"os"
	"strings"
)

func main() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")
	seg.LoadIdf("idf.txt")
	s := bufio.NewScanner(os.Stdin)
	replacer := strings.NewReplacer("，。？?,.;:'\"`~", "")
	for s.Scan() {
		t := s.Text()
		t = replacer.Replace(t)
		for word := range seg.Cut(t, false) {
			if len([]rune(strings.TrimSpace(word))) <= 1 {
				continue
			}
			fmt.Println(word)
		}
	}
}
