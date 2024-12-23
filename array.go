package main

import "fmt"

func main() {
	var hobby [3]string

	hobby[0] = "Bola"
	hobby[1] = "Futsal"
	hobby[2] = "Basket"

	fmt.Println((hobby))

	var values = [...]int{20, 13, 14, 12, 11}
	fmt.Println(len(values))

}
