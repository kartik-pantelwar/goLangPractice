//Differences than C++
//Don't require Paranthases()
//variable declaration not allowed in loop initilization

package main

import (
	"fmt"
)

func main() {
	//for loop

	//for init; condition; post {}

	// the init statement: executed before the first iteration
	// the condition expression: evaluated before every iteration
	// the post statement: executed at the end of every iteration

	// var i int
	// for i=0 ; i<5 ; i++ {
	// 	fmt.Println(i)
	// }

	//simple method, similar to C++

	// for i:=0;i<5;i++ {
	// 	fmt.Println(i)
	// }


	//for loop without init and post statement (short statement)

	// i:= 0
	// for ; i<5 ; {
	// 	fmt.Println(i)
	// 	i++
	// }


	//for each loop
	//we can use it for iterating through array or slice

	// for i:= range object{	//object = array or slice
	// 	fmt.Println(i)
	// }

	//traversing using loops
	//similar to foreach loop
	// for key, value := range map1{ //we can write anything instead of 'key' and 'value'
	// 	fmt.Printf("For Key %v, value is %v\n",key, value)
	// }

	//for loop as while loop
	//semicolans(;) not required

	// i:=0
	// for i<5 {
	// 	fmt.Println(i)
	// 	i++
	// }


	label4:	//we need to put colon after label
		fmt.Println("I am 4")

	//using goto in loop
	//*the label which gets executed first will be the only label
	//*label should be written after the loop
	//*if the label gets executed, then the loop breaks 
	for i:=0;i<6;i++ {
		fmt.Println(i)
		if i==3 {
			goto label3;
		} else if i==4 {
			goto label4
		} else if i==5 {
			goto label5
		}
		
	}



	//using break in loop

	// for i:=0;i<10;i++ {
	// 	if i==4 {
	// 		break;
	// 	}
	// 	fmt.Println(i)
	// }

	//using continue in loop

	// for i:=0;i<10;i++ {
	// 	if i==4 {
	// 		i++	  //any peice of work we can add
	// 		continue;
	// 	}
	// }

	label5:
		fmt.Println("I am 5")
	label3:	//we need to put colon after label
		fmt.Println("I am 3")
	
}
