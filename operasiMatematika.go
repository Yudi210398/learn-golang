package main

import "fmt"

func main() {

	var a = 10
	var b = 20
	var c = 15
	fmt.Println(a + b - c)
	a += b
	fmt.Println(a, b)
	b--
	fmt.Println(a, b)

}
