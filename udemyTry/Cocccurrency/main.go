package main

import (
	"fmt"
	"sync"
	"time"
)
func printSomething(w *sync.WaitGroup, i int, s string){
	time.Sleep(time.Second)
	fmt.Println(i,s)
	defer w.Done()
}
func main(){
	slc:= []string{
		"hi",
		"hello",
		"by",
		"ok",
		"nice",
		"chowmein",
		"pizza",
	}
	var wg sync.WaitGroup
	wg.Add(len(slc))
	for i,val:= range slc{
		
		go	printSomething(&wg,i,val)
	}
	wg.Wait()
}