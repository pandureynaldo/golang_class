package main

import "fmt"

func VariadicFunc(nums...int) {
	// nums := []int{1, 2, 3}
	for key, val := range nums {
		// fmt.Println("Key dari map ", k)
		fmt.Printf("Parameter %d = %d\n", key, val)
	}
}
