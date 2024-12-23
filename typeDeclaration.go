package main

import "fmt"

func main() {
	type Umur int16
	type NoKtp string
	var age Umur = 27
	var ktp NoKtp = "6525454515645156"
	fmt.Println(age, ktp)

}
