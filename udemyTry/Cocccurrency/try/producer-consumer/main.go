package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

//
const NumberOfPizzas =10
var PizzaMade, PizzaFailed, total int
type Producer struct{
	//channel
	data chan PizzaOrder
	quit chan chan error
}
type PizzaOrder struct{
	PizzaNumber int
	Message string
	Success bool
}

func(p *Producer) Close() error{
	ch:= make(chan error)
	p.quit<-ch
	return <-ch
}
func makePizza(pizzaNumber int)*PizzaOrder{
	pizzaNumber++
	if pizzaNumber<=NumberOfPizzas{
		delay:=rand.Intn(5)+1
		fmt.Printf("Received Order number %d\n",pizzaNumber)
		rnd:= rand.Intn(12)+1
		msg:=""
		success:=false
		if rnd<5{
			PizzaFailed++
		}else{
			PizzaMade++
		}
		total++
		fmt.Printf("Making Pizza Number #%d. It will take %d seconds...\n",pizzaNumber,delay)
		time.Sleep(time.Duration(delay)*time.Second)

		if rnd<=2{
			msg = fmt.Sprintf("***We Ran out of ingredients for making pizza number- %d",pizzaNumber)
		}else if rnd<=2{
			msg = fmt.Sprintf("***The Cook quit while making pizza- %d",pizzaNumber)
		}else{
			success=true
			msg = fmt.Sprintf("Pizza number- %d is ready",pizzaNumber)
			pizzaStatus:= PizzaOrder{PizzaNumber: pizzaNumber, Message: msg, Success: success}
			return &pizzaStatus
		}

	}
	return &PizzaOrder{PizzaNumber: pizzaNumber}
}

func PizzaShop(pizzaMaker *Producer){
	//keep track of which pizza we are making
	i:=0
	//run forever or until we receive a quit notification
	//try to make pizza
	for{
		currPizza:= makePizza(i)
		if currPizza!=nil{
			i=currPizza.PizzaNumber
			select{
			case pizzaMaker.data <- *currPizza:
			case quitChan := <-pizzaMaker.quit:
				//close channels
				close(pizzaMaker.data)
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
	color.Cyan("The PizzaShop is open for Business")	
	color.Cyan("-----------------------------------")	

	//create producer 
	pizzaJob:= &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the produdcer in the background
	go PizzaShop(pizzaJob)

	//create and run consumer
	for i:= range pizzaJob.data{
		if i.PizzaNumber<=NumberOfPizzas{
			if i.Success{
				color.Green(i.Message)
				color.Green("Order #%d is out for delivery!",i.PizzaNumber)
			}else{
				color.Red(i.Message)
				color.Red("This Customer is really Mad!")
			}
		} else {
			color.Cyan("Done making Pizza..")
			err:= pizzaJob.Close()
			if err!=nil{
				color.Red("***Error closing Channel!",err)
			}
		}
	}
	//print out the ending message
	color.Cyan("--------------------")
	color.Cyan("Done for the day!!!!")
	color.Cyan("We have successfully made %d pizza, but failed to deliver %d pizzas, with total %d attempts.",PizzaMade, PizzaFailed, total)
	
}