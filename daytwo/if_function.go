package main

import "math/rand"
import "fmt"

func GetRandomNumber(){
	//inisialisasi variabel
	random_number := rand.Intn(20)
	fmt.Println(random_number)
	if random_number < 10 {
		fmt.Println("Nomor lebih kecil dari 10")
	}else if random_number == 10{
		fmt.Println("Nomor sama dengan 10")
	}else{
		fmt.Println("Nomor lebih besar dari 10")
	}
}