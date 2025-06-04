package main
import (
	"fmt"
)
func describe(i interface{}){
	fmt.Printf("(%v, %T)\n",i,i);
}

func main(){
	var i interface{}; //empty interface(can hold value of any type)
	
	i = 25;
	describe(i);

	i="Kartik";
	describe(i);

	i=10.23;
	describe(i)
}