package main

import "fmt"

type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human // anonymous field
    school string
}

type Employee struct {
    Human // anonymous field
    company string
}

// define a method in Human
func (h *Human) SayHi() {
    fmt.Printf("Hi, nama saya %s dan kalian dapat menghubungi saya di no %s\n", h.name, h.phone)
}

// overwrite
func (e *Employee) SayHi() {
    fmt.Printf("Hi, nama saya %s dan umur saya %d tahun\n", e.name, e.age)
}

func main() {
    upin := Employee{Human{"Upin", 22, "123-456-789"}, "Google Inc"}
    ipin := Student{Human{"Ipin", 22, "987-654-321"}, "MIT"}

    upin.SayHi()
    ipin.SayHi()
}