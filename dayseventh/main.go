package main


import "fmt"


type Human struct{
	name string
	age int
	phone string
}

type Student struct{
	Human // anonym field
	school string
}

type Employee struct{
	Human // anonym field
	company string
}


//define a method in Human

func (h *Human) SayHi(){
	fmt.Printf("\nHi, nama saya %s dan nomor telepon saya %s", h.name, h.phone)
}


func main(){
	upin := Employee{Human{"Upin", 21, "08XXX"}, "Hactive"}
	ipin := Student{Human{"Ipin", 20, "081XXXX"}, "Trisakti"}

	upin.SayHi()
	ipin.SayHi()
}