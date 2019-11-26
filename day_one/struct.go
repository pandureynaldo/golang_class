package main

import "fmt"

type person struct {
	name string
	age  int
}

func Struct_1() {

	// var P person

	// P.name = "Pandu"
	// P.age = 25
	// fmt.Println("Nama saya adalah", P.name)

	P := person{"Pandu", 25}
	fmt.Println(P.name)
}
