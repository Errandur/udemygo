package main

import "fmt"

func main () {

	age := 33

	switch age {
		case 16: fmt.Println("Underage!")
		case 17: fmt.Println("Underage!")
		case 18: fmt.Println("You're Good!")
		case 90: fmt.Println("You're Ancient!")
		default:	fmt.Println("Howdy!")
	}

}