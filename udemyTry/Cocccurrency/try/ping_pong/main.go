package main

import (
	"fmt"
	"strings"
)
func shout(ping <-chan string, pong chan<- string){
	//ping is receive only channel
	//pong is send only channel
	for{
		s:= <-ping
		pong<-fmt.Sprintf("%s!!!",strings.ToUpper(s))
	}
}
func main(){
	ping:= make(chan string)
	pong:= make(chan string)

	go shout(ping,pong)

	fmt.Println("Type something and press ENTER (Enter q to QUIT)")
	for{
		//print aprompt
		fmt.Print("-> ")

		var userInput string
		_,_ =fmt.Scanln(&userInput)
		if userInput == strings.ToLower("q"){
			break
		}
		ping<-userInput

		//wait for a resposne

		response:= <-pong
		fmt.Println("Response:",response)
	}
	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
