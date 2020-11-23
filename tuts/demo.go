package main

	import (
			"fmt"
			"time"
			"strconv"
			)

	var pizzaNum = 0
	var pizzaName = ""
		
	func makeDough(stringChan chan string) {
		pizzaNum++
		pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
		fmt.Println("Make dough and send for sauce...")
		stringChan <- pizzaName // Passing pizzaName to the Channel stringChan
		time.Sleep(time.Millisecond *3000)
	}

	func addSauce(stringChan chan string) {
		pizza := <- stringChan // Pulling pizzaName from Channel stringChan
		fmt.Println("Adding sauce and send", pizza, "for toppings...")
		stringChan <- pizzaName // Passing pizzaName to the Channel stringChan
		time.Sleep(time.Millisecond * 3000)
	}

	func addToppings(stringChan chan string) {
		pizza := <- stringChan // Pulling pizzaName from Channel stringChan
		fmt.Println("Adding toppings to", pizza, "and packing up.")
		time.Sleep(time.Millisecond * 3000)
	}
	
	func main() {  
		stringChan := make(chan string)
		fmt.Println("New Order!")
		defer fmt.Println("Order Complete!")
		for i := 0; i < 3; i++ {
			go makeDough(stringChan)
			go addSauce(stringChan)
			go addToppings(stringChan)
			time.Sleep(time.Millisecond * 3000)		
		}
		
	}