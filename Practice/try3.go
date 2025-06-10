package main

import (
	"fmt"
	"sync"
)
func update(w *sync.WaitGroup, result *[]int, data int){
	defer w.Done()
	newData:= data*2
	*result= append(*result, newData)
}
func main(){
	var wg sync.WaitGroup
	input:= []int{1,2,3,4,5}
	result:= []int{}
	for _,data:= range input{
		go update(&wg, &result, data)
	}
	wg.Wait()
}