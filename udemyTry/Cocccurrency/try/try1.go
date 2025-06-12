package main

import (
	"fmt"
	"time"
)

// func add(a, b, c int){
// 	c=a+b;	//c is passed by value,not by reference
// }

func addChannel(a,b int, d chan int){
	d<-a+b;
	// <-d
}

func main(){
	//channels
	a:=1
	b:=2
	// var c int
	// add(a,b,c) //without channel
	// fmt.Println(c)
	d:= make(chan int)
	go addChannel(a,b,d)
	time.Sleep(2*time.Second)
	fmt.Println(<-d)
}