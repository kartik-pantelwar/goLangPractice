package main

import (
	// "fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// data := make([]byte,999999)
	// req,err:= http.Get("http://www.google.com")
	// if err!=nil {
	// 	panic(err)
	// } else{
	// 	req.Body.Read(data)
	// }
	// fmt.Println(string(data))

	//^ smaller approach of doing this
	// res, err := http.Get("http://www.google.com")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	io.Copy(os.Stdout,res.Body)
	// }

	//* Writing Body to file
	file,err := os.Create("data.html")
	res, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err)
	} else {
		io.Copy(file,res.Body)
	}
}
