package main

import (
	"fmt"
	"reflect"
)

func main() {
	println("Arrays e Slices")
	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	fmt.Println(arr1)

	var arr2 = [...]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	slice1 := arr1[:]
	slice2 := arr2[1:3]
	fmt.Println(slice1)
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(arr2))
	fmt.Println(reflect.TypeOf(slice1))

	slice1 = append(slice1, 60)
	fmt.Println(slice1)

	// Array internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacity

}
