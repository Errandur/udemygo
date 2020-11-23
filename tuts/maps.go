package main 

import "fmt"

func main() {
	
	SuperHero := map[string]map[string]string{
		
		"Superman" : map[string]string{
		"Real Name" : "Clark Kent",
		"City" : "Metropolis",
		},

		"Batman" : map[string]string{
			"Real Name" : "Bruce Wayne",
			"City" : "Gotham",
		},

	}

	if temp, hero := SuperHero["Superman"]; hero {
		fmt.Println(temp["Real Name"])
	}
	
}