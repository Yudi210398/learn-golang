package main

import "fmt"

func main() {
	var name string

	name = "Yudi Runat"

	fmt.Println(name, `cak`)

	var fullname = "yudi Runat Masneno"
	fullname = "Yudi Aja"
	fmt.Println(fullname, `cak`)

	binatang := "rusa"
	binatang = "kecoa"
	fmt.Println(binatang, `cak`)

	// multiple variabel

	var (
		nama  = "Yudi Runat"
		marga = "Mansneno"
		agama = "kristen"
	)
	fmt.Println(nama, marga, agama, `cak`)
}
