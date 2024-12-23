package main

import "fmt"

func main() {
	hobby := [...]string{"Sepak Bola", "Dance", "Voli", "Futsal", "Golf"}

	var slice = hobby[2:4]
	slice1 := hobby[:5]
	fmt.Println(slice, slice1)

	hari := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	hariBaru := hari[5:]

	hariBaru[0] = "Sabtu Keren"
	hariBaru[1] = "Minggu Keren"

	haribaru1 := append(hariBaru, "Hari Baru ye")

	fmt.Println(hari, hariBaru, haribaru1)
}
