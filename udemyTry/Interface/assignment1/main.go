package main 
import(
	"fmt"
)

type shape interface{
	GetArea() float64
}

type square struct{
	Side float64
}

type triangle struct{
	Base float64
	Height float64
}

func (s square)GetArea()float64{
	Area:= s.Side*s.Side
	return Area;
}

func(t triangle)GetArea()float64{
	Area:= t.Base * t.Height / 2
	return Area
}

func printArea(s shape)float64{
	return s.GetArea()
}

func main(){
	fmt.Println("hello")
}