package main

import (
	"fmt"
)
func main(){
	arr := [5]int{1: 11, 3: 33} // Elements 0,2,4 will be zero
	for i:= range arr{
		fmt.Println(arr[i])
	}
}
