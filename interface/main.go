package main

import "fmt"

type Human struct{
	name string
	age int
	phone string
}

type Student struct{
	Human // anonymous field
	school string
}

type Employee struct{
	Human // anonymous field
	company string
}

//interface men implemented byby Human, Student and Employee
type Men interface{
	SayHi()
	Sing(title string)
}

func (h Human) SayHi(){
	fmt.Printf("\nHi, nama saya adalah %s dan nomor telepon saya adalah %s", h.name, h.phone)
}

// method sing
func (h Human) Sing(title string) {
    fmt.Println("\nSaya suka menyanyikan lagu ", title)
}


func main(){
	upin := Employee{Human{"Upin", 21, "08XXX"}, "Hactive"}
	ipin := Student{Human{"Ipin", 20, "081XXXX"}, "Trisakti"}

	var iName Men

	//i Name store employee
	iName = upin
	iName.SayHi()
	iName.Sing("I don't love you")

	//iname store student
	iName  = ipin
	iName.SayHi()
	iName.Sing("Kangen")

}
