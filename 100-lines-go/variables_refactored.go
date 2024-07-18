package main


import "fmt"


func main() {
	// `:=` this is called short variable declaration
	a := 42
	b,c := true, 32.0
	d := "string"
	fmt.Println(a, b, c, d)
}
