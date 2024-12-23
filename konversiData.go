package main

import "fmt"

func main() {
	var nilai32 = 32768
	var nilai64 = int64(nilai32)
	var nilai16 = int16(nilai32)

	fmt.Println(nilai32, nilai64, nilai16)

	var name = "Yudi Runat Masneno"
	var getY = name[0]

	fmt.Printf(string(getY), getY)

}
