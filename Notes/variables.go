package main

//var a = 10  *valid globally (outside main)
//a := 10 * not valid globally (inside main only)

const Pi = 3.14 // first capital letter denotes that it is public

import (
	"fmt"
)

func add(x, y int)(sum int){
	return x+y
}

func main(){

	//methods of defining variable in golang

	var a int
	a =10
	b:=20
	var c = 30
	fmt.Println(a+b+c)

	var a int64 = 5
	var b float64 = float64(a)
	fmt.Println(a)
	fmt.Println(b)
}