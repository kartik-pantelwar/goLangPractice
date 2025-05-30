package main

import (
	"fmt"
	"sort"
)
func main(){
	//slices are like Vectors

	//syntax 1 of defining slice -->

	var slice1 =[]int{11,22,33,44}	//if we use this syntax, then we need to initialize it as well
	fmt.Println("slice-",slice1)

	//adding elements using append() method
	//syntax-> slice1= append(slice1,ele1, ele2, ...)
	slice1 = append(slice1, 55,66);
	fmt.Println("slice1 after appending-",slice1)

	//slicing-> slice1= append(slice1[start:end]) 
	slice1 = append(slice1[1:4])
	fmt.Println("after slicing-",slice1)

	//syntax2 of defining slice -->
	//using make()
	//NameOfArray/Slice := make([]dateype, size)
	slice2 := make([]int, 4)
	slice2[0]=3;
	slice2[1]=6;
	slice2[2]=9;
	slice2[3]=12;
	slice2 = append(slice2, 10, 5, 14)
	fmt.Println("slice2=",slice2)
	sort.Ints(slice2)
	fmt.Println("Slice after sorting-",slice2)
}