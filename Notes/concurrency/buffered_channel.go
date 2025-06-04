package main

import (
	"fmt"
)

//unbuffered channel->
//can't store data, immediate receiving required

// func main(){
// 	ch:=make(chan string);
// 	ch<-"hi";
// 	ch<-"hello";
// 	fmt.Println(<-ch);
// 	fmt.Println(<-ch);
// }

//buffered channel->
//can store data, upto the buffer size defined

func main() {
	ch := make(chan string, 2) //*2 is the number of times channel can receive
	//If you try to send more than 2 messages without reading, the program will pause until space is freed.
	ch <- "hi"
	ch <- "hello"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
