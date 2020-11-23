package main 

import "fmt"

func main() {
	
	defer FirstRun()
	SecondRun()
	
}

func FirstRun() {
	fmt.Println("Executed First")
}

func SecondRun() {
	fmt.Println("Executed Second")
}