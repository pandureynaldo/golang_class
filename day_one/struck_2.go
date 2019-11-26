package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	class string
}

func Struct_2() {
	student1 := Student{Human{"Pandu Reynaldo", 25}, "Golang Class"}
	fmt.Println("Nama saya adalah", student1.name)
	fmt.Println("Umur Saya adalah", student1.age)
	fmt.Println("Saya ikut kelas ", student1.class)
}
