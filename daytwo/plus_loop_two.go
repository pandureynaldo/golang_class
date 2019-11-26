package main


import "fmt"

func SumPlusTwo(){
	sum :=0
	for sum < 10{
		sum+=sum
		sum++
	}
	fmt.Println(sum);
}