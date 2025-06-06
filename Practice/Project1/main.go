package main

import (
	"fmt"
	"net/http"
	"sync"
)

func webScrapper(name string, url string, w *sync.WaitGroup){
	defer w.Done()
	res,err:= http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s- %v\n",name,res.Status)
}

func main() {
	//& creating a waitgroup
	var wg sync.WaitGroup;

	wg.Add(4)	//& Adding number of go routines to waitgroup

	//* without slice, one by one
	
	go webScrapper("Google","http://www.google.com",&wg)
	
	go webScrapper("Facebook","http://www.facebook.com",&wg)
		
	go webScrapper("YouTube","http://www.youtube.com",&wg)
	
	// webScrapper("http://www.zomato.com",&wg)

	go webScrapper("Yahoo","http://www.yahoo.com",&wg)

	wg.Wait()	//& Waitgroup waiting for goroutines to execute

	//^with slice, all in one
}
