package main 
import (
	"fmt"
)
func main(){
	ch := make(chan string,3)
	ch<- "messaage 1";
	ch<- "messaage 2";
	ch<- "messaage 3";
	close(ch);
	//channel closed, can't accept any more!

	for msg:= range ch {
		fmt.Println(msg)
	}

}