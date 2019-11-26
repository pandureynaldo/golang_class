package main

import "fmt"

type testInt func(int) bool // definisikan sebuah variabel dengan tipe fungsi 


func FindGreater(param1, param2 int) int {
	if param1 > param2{
		return param1
	}
	return param2
}


func SumAndTime(param1, param2 int) (int,int){
	return param1+param2, param1*param2
}

func VariadicFunc(args...int){
	for key, val := range args{
		fmt.Printf("Parameter index ke %d = %d\n", key, val)
	}
}

func IsOdd(integer int) bool{
	return integer%2 != 0
}

func IsEven(integer int) bool{
	return integer%2==0
}

func filter(dt [] int, fn testInt ) [] int {
	var result [] int
	for _, value := range dt{
		if fn(value){
			result = append(result, value)
		}
	}
	return result
}

// func ReadWriteDefer() bool{
// 	file.Open("filename")
// 	defer file.Close()

// 	if failureX {
// 		return false
// 	}

// 	if failureY {
// 		return false
// 	}

// 	return true
// }

func FmtDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

