package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func testfunc1() {
	a := "❸⒋㈣₄五叁"
	for _, r := range a {
		fmt.Println(unicode.IsNumber(rune(r)))
	}
}

func GetPureContents(a string) string {
	re := regexp.MustCompile("[]*[\ufe00]*[\ufe01]*[\ufe02]*[\ufe03]*[\ufe04]*[\ufe05]*[\ufe06]*[\ufe07]*[\ufe08]*[\ufe09]*[\ufe0a]*[\ufe0b]*[\ufe0c]*[\ufe0d]*[\ufe0e]*[\ufe0f]*[\u1ab0]*[\u1ab1]*[\u1ab2]*[\u1ab3]*[\u1ab4]*[\u1ab5]*[\u1ab6]*[\u1ab7]*[\u1ab8]*[\u1ab9]*[\u1aba]*[\u1abb]*[\u1abc]*[\u1abd]*[\u1abe]*[\u1abf]*[\u1ac0]*[\u1ac1]*[\u1ac2]*[\u1ac3]*[\u1ac4]*[\u1ac5]*[\u1ac6]*[\u1ac7]*[\u1ac8]*[\u1ac9]*[\u1aca]*[\u1acb]*[\u1acc]*[\u1acd]*[\u1ace]*[\u1acf]*[\u1ad0]*[\u1ad1]*[\u1ad2]*[\u1ad3]*[\u1ad4]*[\u1ad5]*[\u1ad6]*[\u1ad7]*[\u1ad8]*[\u1ad9]*[\u1ada]*[\u1adb]*[\u1adc]*[\u1add]*[\u1ade]*[\u1adf]*[\u1ae0]*[\u1ae1]*[\u1ae2]*[\u1ae3]*[\u1ae4]*[\u1ae5]*[\u1ae6]*[\u1ae7]*[\u1ae8]*[\u1ae9]*[\u1aea]*[\u1aeb]*[\u1aec]*[\u1aed]*[\u1aee]*[\u1aef]*[\u1af0]*[\u1af1]*[\u1af2]*[\u1af3]*[\u1af4]*[\u1af5]*[\u1af6]*[\u1af7]*[\u1af8]*[\u1af9]*[\u1afa]*[\u1afb]*[\u1afc]*[\u1afd]*[\u1afe]*[\u1aff]*[\u0300]*[\u0301]*[\u0302]*[\u0303]*[\u0304]*[\u0305]*[\u0306]*[\u0307]*[\u0308]*[\u0309]*[\u030a]*[\u030b]*[\u030c]*[\u030d]*[\u030e]*[\u030f]*[\u0310]*[\u0311]*[\u0312]*[\u0313]*[\u0314]*[\u0315]*[\u0316]*[\u0317]*[\u0318]*[\u0319]*[\u031a]*[\u031b]*[\u031c]*[\u031d]*[\u031e]*[\u031f]*[\u0320]*[\u0321]*[\u0322]*[\u0323]*[\u0324]*[\u0325]*[\u0326]*[\u0327]*[\u0328]*[\u0329]*[\u032a]*[\u032b]*[\u032c]*[\u032d]*[\u032e]*[\u032f]*[\u0330]*[\u0331]*[\u0332]*[\u0333]*[\u0334]*[\u0335]*[\u0336]*[\u0337]*[\u0338]*[\u0339]*[\u033a]*[\u033b]*[\u033c]*[\u033d]*[\u033e]*[\u033f]*[\u0340]*[\u0341]*[\u0342]*[\u0343]*[\u0344]*[\u0345]*[\u0346]*[\u0347]*[\u0348]*[\u0349]*[\u034a]*[\u034b]*[\u034c]*[\u034d]*[\u034e]*[\u034f]*[\u0350]*[\u0351]*[\u0352]*[\u0353]*[\u0354]*[\u0355]*[\u0356]*[\u0357]*[\u0358]*[\u0359]*[\u035a]*[\u035b]*[\u035c]*[\u035d]*[\u035e]*[\u035f]*[\u0360]*[\u0361]*[\u0362]*[\u0363]*[\u0364]*[\u0365]*[\u0366]*[\u0367]*[\u0368]*[\u0369]*[\u036a]*[\u036b]*[\u036c]*[\u036d]*[\u036e]*[\u036f]*[\u1dc0]*[\u1dc1]*[\u1dc2]*[\u1dc3]*[\u1dc4]*[\u1dc5]*[\u1dc6]*[\u1dc7]*[\u1dc8]*[\u1dc9]*[\u1dca]*[\u1dcb]*[\u1dcc]*[\u1dcd]*[\u1dce]*[\u1dcf]*[\u1dd0]*[\u1dd1]*[\u1dd2]*[\u1dd3]*[\u1dd4]*[\u1dd5]*[\u1dd6]*[\u1dd7]*[\u1dd8]*[\u1dd9]*[\u1dda]*[\u1ddb]*[\u1ddc]*[\u1ddd]*[\u1dde]*[\u1ddf]*[\u1de0]*[\u1de1]*[\u1de2]*[\u1de3]*[\u1de4]*[\u1de5]*[\u1de6]*[\u1de7]*[\u1de8]*[\u1de9]*[\u1dea]*[\u1deb]*[\u1dec]*[\u1ded]*[\u1dee]*[\u1def]*[\u1df0]*[\u1df1]*[\u1df2]*[\u1df3]*[\u1df4]*[\u1df5]*[\u1df6]*[\u1df7]*[\u1df8]*[\u1df9]*[\u1dfa]*[\u1dfb]*[\u1dfc]*[\u1dfd]*[\u1dfe]*[\u1dff]*[\u20d0]*[\u20d1]*[\u20d2]*[\u20d3]*[\u20d4]*[\u20d5]*[\u20d6]*[\u20d7]*[\u20d8]*[\u20d9]*[\u20da]*[\u20db]*[\u20dc]*[\u20dd]*[\u20de]*[\u20df]*[\u20e0]*[\u20e1]*[\u20e2]*[\u20e3]*[\u20e4]*[\u20e5]*[\u20e6]*[\u20e7]*[\u20e8]*[\u20e9]*[\u20ea]*[\u20eb]*[\u20ec]*[\u20ed]*[\u20ee]*[\u20ef]*[\u20f0]*[\u20f1]*[\u20f2]*[\u20f3]*[\u20f4]*[\u20f5]*[\u20f6]*[\u20f7]*[\u20f8]*[\u20f9]*[\u20fa]*[\u20fb]*[\u20fc]*[\u20fd]*[\u20fe]*[\u20ff]*[\ufe20]*[\ufe21]*[\ufe22]*[\ufe23]*[\ufe24]*[\ufe25]*[\ufe26]*[\ufe27]*[\ufe28]*[\ufe29]*[\ufe2a]*[\ufe2b]*[\ufe2c]*[\ufe2d]*[\ufe2e]*[\ufe2f]*")
	repl_punct := strings.NewReplacer("-", "", "~", "", "!", "", "@", "", "#", "", "$", "", "%", "", "^", "", "&", "", "*", "", "(", "", ")", "", "_", "", "+", "", "{", "", "}", "", "|", "", ":", "", "<", "", ">", "", "?", "", "`", "", "[", "", "]", "", ";", "", "'", "", ",", "", ".", "", "/", "", "～", "", "！", "", "#", "", "…", "", "…", "", "（", "", "）", "", "「", "", "」", "", "、", "", "【", "", "】", "", "；", "", "‘", "", "：", "", "“", "", "”", "", "《", "", "》", "", "？", "", "\"", "", "\\", "")
	repl_digit := strings.NewReplacer("⓪", "0", "①", "1", "②", "2", "③", "3", "④", "4", "⑤", "5", "⑥", "6", "⑦", "7", "⑧", "8", "⑨", "9", "⑩", "10", "⑪", "11", "⑫", "12", "⑬", "13", "⑭", "14", "⑮", "15", "⑯", "16", "⑰", "17", "⑱", "18", "⑲", "19", "⑳", "20", "⑴", "1", "⑵", "2", "⑶", "3", "⑷", "4", "⑸", "5", "⑹", "6", "⑺", "7", "⑻", "8", "⑼", "9", "⑽", "10", "⑾", "11", "⑿", "12", "⒀", "13", "⒁", "14", "⒂", "15", "⒃", "16", "⒄", "17", "⒅", "18", "⒆", "19", "⒇", "20", "⒈", "1", "⒉", "2", "⒊", "3", "⒋", "4", "⒌", "5", "⒍", "6", "⒎", "7", "⒏", "8", "⒐", "9", "⒑", "10", "⒒", "11", "⒓", "12", "⒔", "13", "⒕", "14", "⒖", "15", "⒗", "16", "⒘", "17", "⒙", "18", "⒚", "19", "⒛", "20", "❶", "1", "❷", "2", "❸", "3", "❹", "4", "❺", "5", "❻", "6", "❼", "7", "❽", "8", "❾", "9", "❿", "10", "➊", "1", "➋", "2", "➌", "3", "➍", "4", "➎", "5", "➏", "6", "➐", "7", "➑", "8", "➒", "9", "➓", "10", "⓿", "0", "⓵", "1", "⓶", "2", "⓷", "3", "⓸", "4", "⓹", "5", "⓺", "6", "⓻", "7", "⓼", "8", "⓽", "9", "⓾", "10", "⓫", "11", "⓬", "12", "⓭", "13", "⓮", "14", "⓯", "15", "⓰", "16", "⓱", "17", "⓲", "18", "⓳", "19", "⓴", "20", "㈠", "1", "㈡", "2", "㈢", "3", "㈣", "4", "㈤", "5", "㈥", "6", "㈦", "7", "㈧", "8", "㈨", "9", "㈩", "10", "０", "0", "１", "1", "２", "2", "３", "3", "４", "4", "５", "5", "６", "6", "７", "7", "８", "8", "９", "9", "⁰", "0", "¹", "1", "²", "2", "³", "3", "⁴", "4", "⁵", "5", "⁶", "6", "⁷", "7", "⁸", "8", "⁹", "9", "₀", "0", "₁", "1", "₂", "2", "₃", "3", "₄", "4", "₅", "5", "₆", "6", "₇", "7", "₈", "8", "₉", "9", "零", "0", "一", "1", "二", "2", "三", "3", "四", "4", "五", "5", "六", "6", "七", "7", "八", "8", "九", "9", "〇", "0", "壹", "1", "贰", "2", "叁", "3", "肆", "4", "伍", "5", "陆", "6", "柒", "7", "捌", "8", "玖", "9", "0️", "0", "1️", "1", "2️", "2", "3️", "3", "4️", "4", "5️", "5", "6️", "6", "7️", "7", "8️", "8", "9️", "9")
	ret := re.ReplaceAllString(a, "")
	ret = repl_digit.Replace(ret)
	ret = repl_punct.Replace(ret)
	return ret
}

func main() {
	testfunc1()
	a := "你好啊我是③⑹⒉❽➓⓻㈣７⁸₃六壹1⃣️3⃣️"
	fmt.Println(GetPureContents(a))
}
