package main

import "fmt"


func testFunc() {
    number := 0

	testLabel:
		fmt.Println(number)
		number++
		goto testLabel
}