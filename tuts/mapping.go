package main

import (
	"fmt"
)

func main() {
	personSalary := make(map[string]int)
	personSalary["Ethan"] = 9999
	personSalary["Cooper"] = 10000
	personSalary["Mason"] = 12000
	fmt.Println("personSalary map contents:", personSalary)

	for key, value := range personSalary {
		fmt.Println("Key:", key, "Value:", value)
	}
}
