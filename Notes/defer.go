//defer executes a statement after the whole function is completed
//No matter in which seequence the statement is written
//It is used in opening and closing files from os

//*Follows LIFO Principle

//ex->
//func openFile(){ {os.open(filename);
//defer file.close(os);


package main
import (
	"fmt"
)
func myfunc(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}
func main(){
	defer fmt.Println("End");
	defer fmt.Println("Mid");
	fmt.Println("Start");
	myfunc();

}