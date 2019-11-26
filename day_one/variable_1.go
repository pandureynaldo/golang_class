package main

import "fmt"

func VariableGo1() {
	/* Pembubuhan value 1 baris */
	// var nama string = "Pandu Reynaldo"
	// fmt.Println(nama)

	/* Pembubhan value baris berikutnya */
	// var umur int
	// umur = 25

	// fmt.Println(umur)

	// alamat := "Jakarta"
	// fmt.Println(alamat)

	var (
		nama   = "Pandu Reynaldo"
		umur   = 17
		alamat = "Indramayu"
	)
	fmt.Printf("%s berumur %d tahun\n", nama, umur)
	fmt.Println("Beralamat di", alamat)

}
