package main

import "fmt"

func main() {

	hobby := [...]string{"Sepak Bola", "Dance", "Voli", "Futsal", "Golf"}

	// ! pada slice kalo mangambil slice awal dimulai dari index 0 tapi kalo di akhir dimulai dari index 1 atau pertama [2:4 ]
	var slice = hobby[2:4]
	slice1 := hobby[:5]
	fmt.Println(slice, slice1, `hasil`)

	hari := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	hariBaru := hari[5:]

	hariBaru[0] = "Sabtu Keren"
	hariBaru[1] = "Minggu Keren"

	haribaru1 := append(hariBaru, "Hari Baru ye")
	haribaru1[0] = `Sabtu lama`
	hariBaru[0] = `Sabtu cool`
	fmt.Println(hari)
	fmt.Println(hariBaru)
	fmt.Println(haribaru1)

	newsliceMake := make([]string, 2, 5)

	newsliceMake[0] = "natan"
	newsliceMake[1] = ",iftah"
	newsliceMake1 := append(newsliceMake, `sasa`, `Gokil`, `Datas`, `Haris`, `Woy `, `dasd`)
	newsliceMake1[6] = `lika liku`
	fmt.Println(newsliceMake)
	fmt.Println(newsliceMake1)

	// ! copy Array Slice

	fromCopyArray := hari[:]

	toCopyArray := make([]string, len(fromCopyArray), cap(fromCopyArray))

	copy(toCopyArray, fromCopyArray)

	fmt.Println(`kita copy`, toCopyArray, fromCopyArray)
}
