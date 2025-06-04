// package main

// import "fmt"

// // Function that sends data into a channel
// func sendData(ch chan string) {
//     ch <- "Hello from Goroutine!" // Sends data
// }

// func main() {
//     ch := make(chan string) // Create an unbuffered channel

//     go sendData(ch) // Start goroutine

//     message := <-ch // Receive data from channel
//     fmt.Println(message) // Output: Hello from Goroutine!
// }


package main
import (
	"fmt"
)
func sendData(ch chan string, Data string){
	ch<-Data;
	fmt.Println("Data received Sucessfuly")
}
func fetchData(ch chan string)string {
	RData:= <-ch;
	return RData;
}
func main(){
	//sending Data
	channel1 := make(chan string);
	Data := "Important Channel Data, kisi ko pta nhi chalna chahiye";
	go sendData(channel1, Data);
	//*go is used because it is an unbuffered channel. with the use of 'go'.
	//Without go, the function will run synchronously and block until it completes. It can safely send data without errors, or causing any deadclocks, and without the need of instant receiver.
	//It won't work without  go

	//receiving Data
	milaData:= fetchData(channel1);
	fmt.Println("*Data Received=",milaData)
	
}


