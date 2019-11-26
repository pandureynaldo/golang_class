package main

import "fmt"

func GetRangeMap(){
	nums := [] int  {1,2,3};
	for k, num := range  nums{
		// fmt.Println("Key dari map ", k)
		fmt.Println("Key ", k, " Map ", num)
	}
}