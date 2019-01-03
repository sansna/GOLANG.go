package main
import "strings"
import "fmt"
func main() {
	if strings.HasPrefix("record_rds_all_score_v2", "record_rds_all_score") {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
