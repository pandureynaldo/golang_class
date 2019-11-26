package main

import "fmt"

// import (
//     . "fmt"
// )
// Dot operator. Operator ini digunakan agar kalian tidak perlu menyisipkan nama package lagi ketika menggunakan function pada package yang di-import.
// import (
//     f "fmt"
// )

//Alias operator. Operator ini digunakan untuk mengganti nama package yang di-import dengan nama yang lebih mudah atau lebih simple.

func main(){
	no1 := 7
	no2 := 2
	fmt.Printf("max between %d - %d is %d", no1, no2, FindGreater(no1,no2))

	no1PLUSno2, no1TIMESno2 := SumAndTime(no1, no2)

	fmt.Printf("\n%d + %d = %d", no1, no2, no1PLUSno2)
	fmt.Printf("\n%d * %d = %d", no1, no2, no1TIMESno2)

	// VariadicFunc(4, 11, 23, 7, 33, 10, 5)

	var data = [] int {1, 2, 3, 4, 5, 6, 7}

	odd := filter(data, IsOdd)
    even := filter(data, IsEven)

    fmt.Println("\nData = ", data)
    fmt.Println("Data ganjil = ", odd)
	fmt.Println("Data genap = ", even)
	
	// if ReadWriteDefer(){
	// 	fmt.Println("Success")
	// }else{
	// 	fmt.Println("Failure")
	// }

	FmtDefer()



}