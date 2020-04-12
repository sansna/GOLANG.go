package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("你好|上线|在吗|在|哈喽|侬好|hello|hi|啊哈|嘎哈|在干嘛|问你|出现")
	s := "对象sdfd.你没啊"
	fmt.Println(re.MatchString(s))
	re = regexp.MustCompile("你.*对象|对象.*你")
	fmt.Println(re.MatchString(s))

	s = "（（、$；、（、（、！+、（、（+、"
	//de_punct := regexp.MustCompile("[~,./<>?()~\\]\\[-\\\\\"'`]+")
	de_punct := regexp.MustCompile("[- !@#$%^&*()\\[\\]=\\+{}\\\\|;':\",./<>?`~～·！@#¥%…（）—【】「」、｜：“”；‘’，。《》/？]")
	var dest string
	dest = de_punct.ReplaceAllString(s, "")
	fmt.Println(dest)
}
