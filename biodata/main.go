package main


import (
	"fmt"
	"os"
	"strconv"
)


type classmate struct{
	name string
	address string
	job string
}

var allClassMate  = [3] classmate{classmate{"Arif","Kalimalang","Backend"},classmate{"Fulan","Kebon jeruk","Admin"}, classmate{"Fulin", "Mampang", "Auditor" }}
var length_array int = 3
func main(){
	
	var index,_ = strconv.Atoi(os.Args[1])
	// getClassmate(index)
	// fmt.Println(index)
	// fmt.Println(length_array)
	index = index - 1
	getClassmate(index)
	

}

func getClassmate(index int){
	if index >= length_array {
		fmt.Println("Tidak ada teman di index ke-", (index+1))
	}else{
		fmt.Println("Berikut adalah inputan anda ", index+1, ":")
		fmt.Println("Nama \t\t: ", allClassMate[index].name)
		fmt.Println("Alamat \t\t: ", allClassMate[index].address)
		fmt.Println("Pekerjaan \t: ", allClassMate[index].job)
	}
}
