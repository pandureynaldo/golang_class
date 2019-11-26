package main

import "fmt"

func CheckNumberIntegerAll(index int){
	switch index{
		case 1:
			fmt.Println("index sama dengan 1")
			fallthrough
		case 2,3,4:
			fmt.Println("index sama dengan 2 atau 3 atau 4")
			fallthrough
		case 10 :
			fmt.Println("index sama dengan 10")
			fallthrough
		default:
			fmt.Println("index adalah integer")
			// fallthrough
	}
}