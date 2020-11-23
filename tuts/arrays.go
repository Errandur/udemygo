package main 

import "fmt"

func main() {
	
	var evenNum[5] int 

	evenNum[0] = 5
	evenNum[1] = 4
	evenNum[2] = 3
	evenNum[3] = 2
	evenNum[4] = 1
	
	fmt.Println(evenNum[4])

	oddNum := [5]int{5,4,3,2,1}
	fmt.Println(oddNum[4])

	for _, value := range evenNum {
		fmt.Println(value)
	}

	numSlice := []int{9,8,7,6,5}

	sliced := numSlice[3:5]
	fmt.Println(sliced)
}