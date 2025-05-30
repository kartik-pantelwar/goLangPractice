package main

import (
	"fmt"
	"io"
	"ioutil"
	"os"
)

//Creating new File and Entering Data into it
func main(){
	data := "This is the new File!"
	file,err := os.Create("./newfile.txt");
	fmt.Println("Creating File..")

	length, err := io.WriteString(file,data)

	if err!=nil {
		panic(err)
	} else{
		fmt.Println("File Created of length-",length)
	}
	defer file.Close();


	//Reading Data from File
	readFile("./newfile.txt")
}

//Reading Data
func readFile(filename string){
	daata, err := ioutil.ReadFile(filename)
	if err!= nil {
		panic(err)
	}
	fmt.Println(daata)
}