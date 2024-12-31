package main

import "fmt"

func main() {
	exapmpleMap := map[string]string{
		"nama":  "Hana Safira",
		"hobby": "Berenang",
	}

	fmt.Print(exapmpleMap, len(exapmpleMap))

	buatMap := make(map[string]string)

	buatMap["Kesukaan"] = "Nasi Goreng"
	buatMap["Emosi"] = "Always Happy"
	buatMap["salah"] = "gk Happy"
	fmt.Println(buatMap, `caks`)
	delete(buatMap, "salah")
	fmt.Println(buatMap, `caks`)

}
