package main

import "fmt"

func main() {
	var array1 [3]int8
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	fmt.Println(array1)

	array2 := [3]int{7, 8, 9}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{0, 0, 1, 1}
	fmt.Println(slice)

	slice = append(slice, 100)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 0
	fmt.Println(slice2)
}
