package main

import "fmt"

func main() {  
		type Employee struct {  
		firstName string
		lastName  string
		age       int
	}
	
	emp1 := Employee{
        firstName: "Ethan",
		lastName: "Shoultz",
        age:       26,
    }
	
	fmt.Println(emp1)
}