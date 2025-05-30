package main 
import (
	"fmt"
)

//*Function Overloading not supported in Go

//For adding 2 numbers
func sum2(a int, b int)(int){
	return a + b
}

func allSum(values ...int)int{ //values can be any name
	total :=0
	for _,val := range values {
		total = total + val
	}
	return total
}

// func (){
// 	//anonymous function
// 	//immediately invoked
// 	fmt.Println("I am annonymous function")
// }()

func swapInt(a int, b int)(int ,int){
	return b,a // we can return multiple things
}

func main(){

	//swapping variables
	// var a int = 10
	// var b int = 20
	// fmt.Println("Before Swap->")
	// fmt.Printf("a- %v, b- %v\n",a,b)
	// a,b = swapInt(a,b);
	// fmt.Println("After Swap->")
	// fmt.Printf("a- %v, b- %v\n",a,b)

	fmt.Println(allSum(10,20,30,40))
	fmt.Println(sum2(10,20))

}