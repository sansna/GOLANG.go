package main
type aa struct {
	a int
}

func main() {
	for {
		a := aa{
			a: 1,
		}
		a = a
	}
	return
}
