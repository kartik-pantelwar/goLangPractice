package main

import (
	"fmt"
)

func main() {
	//declaring array using method 1
	var arr [4]int

	//1st method of printing
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])

	//2nd method of printing
	fmt.Println("my array is-", arr)

	var arr1 = [4]int{11, 22, 33, 44}

	//3rd method of printing array
	i := 0
	for i = 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

}
