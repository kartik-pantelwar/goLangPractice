package main

import (
	"fmt"
	"sync"
)
var SMS string
var wg sync.WaitGroup
func updateSMS(s string){
	SMS = s
	defer wg.Done()
}
func main(){
	wg.Add(2)
	go updateSMS("Hello World")
	go updateSMS("Bye World")
	wg.Wait()
	fmt.Println(SMS)
}

