package main

import "net/http"
import "encoding/json"
import "fmt"

type Kev struct {
	ADD string `json:"t"`
	DFS string `json:"a"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/life", szs)
	http.ListenAndServe(":8888", mux)
}
func szs(w http.ResponseWriter, r *http.Request) {
	var t Kev
	dcd := json.NewDecoder(r.Body)
	dcd.Decode(&t)
	w.Write([]byte("hello world!"))
	fmt.Println(t)
}
