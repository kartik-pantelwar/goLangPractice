package main
import "fmt"

//declaring deck type 
type deck []string;

func (d deck) print(){
	for i,e := range d {
		fmt.Println(i,"=",e)
	}
}
func NewDeck()deck{
	d := deck{}
	suits := []string{"club","diamond","heart","spade"}
	for _,s := range suits{
		for v:=2;v<11;v++ {
			d=append(d, fmt.Sprintf("%s-%d",s,v))
		}
		faceCards := []string{"king","queen","jack","ace"}
		for _,f:= range faceCards {
			d=append(d, fmt.Sprintf("%s-%s",s,f))
		}
	}
	return d;
}
