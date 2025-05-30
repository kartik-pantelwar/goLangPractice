package main

import (
	"fmt"
)

func main() {
	//declaring array using method 1
	var arr [4]int

	//delaring array with literal
	var arr1 = [4]int{11,22,33,44}

	//shorthand declaration
	arr2 := [4]int{1,2,3,4}

	//declaration with automatic size
	arr3 := [...]int{34,523,643,252}

	//1st method of printing
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])

	//2nd method of printing
	fmt.Println("my array is-", arr)

	var arr4 = [4]int{11, 22, 33, 44}

	arr5 := [5]int{1: 10, 3: 30} // Elements 0,2,4 will be zero

	//3rd method of printing array
	i := 0
	for i = 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

}
