package main

import (
	"fmt"
	"net/http"
)
func main(){
	res, err:= http.Get("http://www.youtube.com")	
	if err!=nil{
		panic(err)
	}
	fmt.Println("res=",res.StatusCode)
}