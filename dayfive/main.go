package main

import (
	"fmt"
	"math"
)


type Circle struct{
	radius float64
}

type Rectangle struct{
	width, height float64
}

func areaRectangle(r Rectangle) float64 {
	return r.width * r.height
}

func areaCircle(c Circle) float64{
	return c.radius * c.radius * math.Pi
}

func main(){
	r1 := Rectangle{12,2}
	c1 := Circle{25}
	fmt.Println("Area Rectangle :", areaRectangle(r1))
	fmt.Println("Area Circle :", areaCircle(c1))
}