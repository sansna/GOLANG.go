package main

import (
	"fmt"
	"github.com/wangbin/jiebago/analyse"
)

func main() {
	s := "美食猫北京广州她讲明天我说卧室我是一个没有地球狗的人"
	var te analyse.TagExtracter
	err := te.LoadDictionary("/Users/zuiyou/GO/conf/hanabi-bizsrv-acnt/etc/dict.txt")
	if err != nil {
		fmt.Println(err)
	}
	err = te.LoadIdf("/Users/zuiyou/GO/conf/hanabi-bizsrv-acnt/etc/dict.txt")
	if err != nil {
		fmt.Println(err)
	}
	//te.LoadIdf("idf.txt")
	//te.LoadStopWords("/Users/zuiyou/GO/conf/hanabi-bizsrv-acnt/etc/stopwords.txt")
	zz := te.ExtractTags(s, 10)
	for _, z := range zz {
		fmt.Println(z.Text())
	}
}
