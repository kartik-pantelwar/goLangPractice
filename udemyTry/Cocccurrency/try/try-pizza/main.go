//^ Changes than the orignal program->
// 1. It won't wait if the ingredients are over, or the cook left the job(Immediately tell to the Counter). It will take delay time only when pizza is successfully making
// 2. Removed the unnecessary Close()

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

//constants
const NumberOfPizzas =10
var PizzaMade, PizzaFailed, Total int

//types
type PizzaOrder struct{
	PizzaNumber int
	Message string
	Success bool
}

type Producer struct{
	data chan PizzaOrder
	quit chan chan error
}

//close method
func (p *Producer) Close() error{
	ch:= make(chan error)
	p.quit<-ch
	return <-ch
}

//functions
func makePizza(pizzaNumber int)*PizzaOrder{
	pizzaNumber++
	if pizzaNumber<=NumberOfPizzas{
		success:=false
		msg:=""
		delay:= rand.Intn(5)+1
		rnd:= rand.Intn(12)+1		
		fmt.Printf("Received Order #%d\n",pizzaNumber)
		if rnd<5{
			PizzaFailed++;
		}else{
			PizzaMade++;
		}
		Total++;
		if rnd<=2{
			msg = fmt.Sprintf("We are Out of Ingredients!!!\n")
		}else if rnd<5{
			msg = fmt.Sprintf("Cook left the Job!!!\n")
		}else{
			fmt.Printf("Making Pizza #%d,It will take %d seconds\n",pizzaNumber,delay)
			time.Sleep(time.Duration(delay)*time.Second)
			msg=fmt.Sprintf("Order Number #%d is ready\n",pizzaNumber)
			success = true
		}
		return &PizzaOrder{PizzaNumber: pizzaNumber, Message: msg, Success: success}
	}
	return &PizzaOrder{PizzaNumber: pizzaNumber}
}

//pizzaShop function
func pizzaShop(p *Producer){
	i:=0
	for{
		currPizza:= makePizza(i)
		if currPizza!=nil{
			i=currPizza.PizzaNumber
			select{
			case p.data<-*currPizza:
			case quitChan:= <-p.quit:
				close(p.data)
				close(quitChan)
				return
			}
		}
	}
}

func main(){
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//print starting message
	color.Cyan("Welcome to Pizza Shop...\n")
	color.Cyan("--------------------------\n")

	//create producer(chef)
	Chef:= &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//keep running the producer in background
	go pizzaShop(Chef)

	//create and run consumer(Counter wale)
	for i:= range Chef.data{
		if i.PizzaNumber<=NumberOfPizzas{
			if i.Success{
				color.Green("Pizza #%d is ready for delivery\n",i.PizzaNumber)
			}else{
				color.Red(i.Message)
				color.Red("Customer is Mad for order %d\n",i.PizzaNumber)
			}
		}else{
			err:=Chef.Close()
			if err!=nil{
				color.Red("Panic closing channel ",err)
			}
		}
	}

	//print out ending message
	color.Cyan("Pizza Shop has been Closed...\n")
	color.Cyan("-----------------------------\n")
}