package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Philosopher struct {
	Name  string
	Left  int
	Right int
}

// Philosophers
var philosophers = []Philosopher{
	{Name: "Kartik", Left: 4, Right: 0},
	{Name: "Motu", Left: 0, Right: 1},
	{Name: "Patlu", Left: 1, Right: 2},
	{Name: "Jhatka", Left: 2, Right: 3},
	{Name: "Ghasita", Left: 3, Right: 4},
}

func dine(){
	wg := &sync.WaitGroup{}	//jb sbne kha lia toh 0 hoga
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}	//jb sare table pe baith chuke hai tb 0 hoga, ab taiyar hai khane ke liye
	seated.Add(len(philosophers))

	//maps of fork, telling which map is locked or unlocked to use
	var forks= make(map[int]*sync.Mutex)
	for i:=0;i<len(philosophers);i++{
		forks[i]=&sync.Mutex{}
	}

	//start the meal.
	for i:=0;i<len(philosophers);i++{
		//use a goroutine for current philosopher
		go diningProblem(philosophers[i],wg,forks,seated)
	}
	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup){
	defer wg.Wait()

	//seat the philosopher at the table
	color.Green("%s is seated at the table.\n",philosopher.Name)
	seated.Done()

	seated.Wait()

	//eat three times
	for i:=0;i<hunger;i++{
		forks[philosopher.Left].Lock()
		fmt.Printf("%s has taken the left fork\n",philosopher.Name)
		forks[philosopher.Right].Lock()
		fmt.Printf("%s has taken the Right fork\n",philosopher.Name)

		color.Yellow("Philosopher %s has all the forks, and is eating\n",philosopher.Name)
		time.Sleep(eatTime)

		fmt.Printf("Philosopher %s is thinking\n",philosopher.Name)
		time.Sleep(thinkTime)

		fmt.Printf("Philosopher %s is sleeping\n",philosopher.Name)
		time.Sleep(sleepTime)

		forks[philosopher.Left].Unlock()
		forks[philosopher.Right].Unlock()

		color.Red("%s put down the forks",philosopher.Name)
	}

	fmt.Println(philosopher.Name,"is staisfied")
	fmt.Println(philosopher.Name,"left the table")
}

// func think(){}

// func sleep(){}

// define some variables
var hunger = 3 //how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 3* time.Second
var sleepTime = 1* time.Second
func main() {
	//print start message
	fmt.Println("Dinning Philosophers problem")
	fmt.Println("----------------------------")
	fmt.Println("The Table is empty")	//sare sath mei khayenge, respect hai, koi akela nhi khayega

	//start the meal
	dine()


	//print finish message
	fmt.Println("The Table is empty")
}