package main

import "fmt"

func main() {
	name := "ker"

	if name == "keren" {
		fmt.Println("Masuk bg")
	} else if name == "Good" {
		fmt.Println("keren lu")
	} else {
		fmt.Println("gk keren lu")
	}

	// if statment short

	if kataBukan3 := len(name); kataBukan3 == 3 {
		fmt.Println(` ya pasti katanya 3 `)
	}

}
