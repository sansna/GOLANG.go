package main

//import "strings"
import "fmt"
import "regexp"

func main() {
	a := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	reg := regexp.MustCompile("[-*_壱弐参⓪①②③④⑤⑥⑦⑧⑨⑩⑪⑫⑬⑭⑮⑯⑰⑱⑲⑳⑴⑵⑶⑷⑸⑹⑺⑻⑼⑽⑾⑿⒀⒁⒂⒃⒄⒅⒆⒇⒈⒉⒊⒋⒌⒍⒎⒏⒐⒑⒒⒓⒔⒕⒖⒗⒘⒙⒚⒛❶❷❸❹❺❻❼❽❾❿➊➋➌➍➎➏➐➑➒➓⓿⓵⓶⓷⓸⓹⓺⓻⓼⓽⓾⓫⓬⓭⓮⓯⓰⓱⓲⓳⓴㈠㈡㈢㈣㈤㈥㈦㈧㈨㈩０１２３４５６７８９⁰¹²³⁴⁵⁶⁷⁸⁹₀₁₂₃₄₅₆₇₈₉零一二三四五六七八九〇壹贰叁肆伍陆柒捌玖]")
	//reg := regexp.MustCompile("[-壱弐参⓪①②③④⑤⑥⑦⑧⑨⑩⑪⑫⑬⑭⑮⑯⑰⑱⑲⑳0⑴⑵⑶⑷⑸⑹⑺⑻⑼⑽⑾⑿⒀⒁⒂⒃⒄⒅⒆⒇0⒈⒉⒊⒋⒌⒍⒎⒏⒐⒑⒒⒓⒔⒕⒖⒗⒘⒙⒚⒛0❶❷❸❹❺❻❼❽❾❿0➊➋➌➍➎➏➐➑➒➓⓿⓵⓶⓷⓸⓹⓺⓻⓼⓽⓾⓫⓬⓭⓮⓯⓰⓱⓲⓳⓴0㈠㈡㈢㈣㈤㈥㈦㈧㈨㈩０１２３４５６７８９⁰¹²³⁴⁵⁶⁷⁸⁹₀₁₂₃₄₅₆₇₈₉零一二三四五六七八九〇壹贰叁肆伍陆柒捌玖1️2️3️4️5️6️7️8️9️0️🔟1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣]")
	b := reg.ReplaceAllString(a, "")
	fmt.Println(a, b)

	lena := float64(len([]rune(a)))
	lenb := float64(len([]rune(b)))
	fmt.Println(lena, lenb)

	fmt.Println((lena - lenb) / lena)
}
