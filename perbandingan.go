package main

import "fmt"

func main() {
	var nilaiBudi = 78
	var absenBudi = 88
	var minumumAbsen = 80
	var minumumNiali = 75

	var kelulusan = nilaiBudi > minumumNiali && absenBudi > minumumAbsen

	fmt.Print(kelulusan)

}
