package main
import (
	"fmt"
)

func main(){
	var slice1 = []int{33,64,21,34}
	fmt.Println("slice1-",slice1);
	
	//Removing element from slice from specific index
	//slice= append(slice[:index],slice[index+1:]...)

	//suppose we want to remove 1st index from slice
	slice1 = append(slice1[:1],slice1[2:]...)
	fmt.Println(slice1)	

	// '...' are veriodic functoions 
}