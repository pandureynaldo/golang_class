package main

import "fmt"

func CheckNumberInteger(index int){
	switch index{
		case 1:
			fmt.Println("index sama dengan 1")
		case 2,3,4:
			fmt.Println("index sama dengan 2 atau 3 atau 4")
		case 10 :
			fmt.Println("index sama dengan 10")
		default:
			fmt.Println("index adalah integer")
	}
}