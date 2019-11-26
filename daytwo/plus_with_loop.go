package main

//penggunaan expression
import "fmt"


func SumPlus(){
	sum := 0
	for index := 0 ; index < 10; index++{
		sum+=index;
	}
	fmt.Println(sum)
}