package main
import "fmt"
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x < 10 {
        return true
    }
    a := []int{}
    for ;x != 0; x/=10 {
        a = append(a, x%10)
    }
	fmt.Println(a, len(a)/2-1, a[0:0])
    for idx, i := range a[0:len(a)/2] {
        if i != a[len(a)-1-idx] {
            return false
        }
    }
    return true
}
func main() {
	fmt.Println(isPalindrome(101))
}
