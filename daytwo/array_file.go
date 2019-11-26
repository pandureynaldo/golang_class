package main

import "fmt"

func GetArrayOne(){
	var arrayTest [10] int
	//set data
	arrayTest[0] = 5
	arrayTest[1] = 98

	//get data
	fmt.Printf("Data pertama atau index 0 %d\n", arrayTest[0])
	fmt.Printf("Data pertama atau index 9 %d\n", arrayTest[9])
}

func GetArrayTwo(){
	arrayTest := [4] int {74,12,22,49}

	for index, value := range arrayTest {
		fmt.Printf("Index %d = %d\n", index, value)
	}
}

func GetArrayThree(){
	arrayTest := [2][4] int {[4]int{1,2,3,4}, [4]int{5, 6, 7, 8}}
	arrayTest2 := [2][4] int {{9,8,7,6},{1,2,3,4}}

	for index, value := range arrayTest {
		fmt.Printf("Index %d = %d\n", index, value)
	}

	for index, value := range arrayTest2 {
		fmt.Printf("Index %d = %d\n", index, value)
	}
}

func GetArraySlice(){
	// slice := [] int {1,4,2,3,76}
	data := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var aSlice, bSlice []byte

	// operasi standar
	aSlice = data[:3] // sama dengan aSlice = data[0:3], hasilnya a,b,c
	fmt.Println(string(aSlice))
	aSlice = data[5:] // sama dengan aSlice = data[5:10], 10 merupakan panjang dari data, hasilnya f,g,h,i,j
	fmt.Println(string(aSlice))
	aSlice = data[:] // sama dengan aSlice = data[0:10], hasilnya menampikan semua elemen
	fmt.Println(string(aSlice))

	// re-slice
	aSlice = data[3:7] // mulai dari index=3 sampai kapasitas=7, hasilnya d,e,f,g
	fmt.Println(aSlice)
	bSlice = aSlice[1:3] // mulai dari index=1 sampai kapasitas=3, hasilnya e,f
	fmt.Println(bSlice)
	bSlice = aSlice[:3] // mulai dari index=0 sampai kapasitas=3, hasilnya d,e,f
	fmt.Println(bSlice)
	bSlice = aSlice[0:5] // slice dapat bertambah melebihi kapasitasnya karena slice hanya mng-copy data (nanti dipelajari lebih lanjut), hasilnya d,e,f,g,h
	fmt.Println(bSlice)
	bSlice = aSlice[:] // bSlice memiliki elemen seperti aSlice, hasilnya d,e,f,g
	fmt.Println(bSlice)
}

func GetArraySliceMake(){
	slice := make([]int,0, 10)
	slice = append(slice,8)
	fmt.Println(slice)

	slice = slice[0:8]
	slice[7] = 9
	fmt.Println(slice);

	scores := []int{1, 2, 3, 4, 5}
	slice = scores[2:4]
	// slice[0] = 999

	fmt.Println("slice: ", slice)
	fmt.Println("scores: ", scores)
}

func GetArraySliceCopy(){
	scores := [] int {1,2,3,4,5}
	newScores := make([]int, 5)
	copy(newScores, scores)
	newScores[0] = 999
	scores[1] = 1000

	fmt.Println("copy: ", newScores)
	fmt.Println("scores: ", scores)
}

func GetMapArray(){
	// var numbers map[string]int
	numbers_two := make(map[string]int)

	// numbers["one"] = 1
	numbers_two["twenty"] = 21

	// fmt.Println(numbers["one"])
	fmt.Println(numbers_two["twenty"])
}

func GetMapAdvance(){
	// initialize map with key and value
	rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2}

	// map memiliki 2 return value, jika key tersedia pada map maka 'exist' akan true tetapi bila tidak ada maka akan false
	// cSharpRating, exist := rating["C#"]
	for index, _ := range rating{
		exist := rating[index] 
		// if rating[index] {
		// 	fmt.Printf(" %s tersedia pada map dan memiliki %f ", index,value)
		// } else {
		// 	fmt.Printf(" %s tersedia pada map tidak memiliki %f ", index,value)
		// }
		fmt.Println(exist);
	} 

	// delete elemen
	delete(rating, "C")
}


