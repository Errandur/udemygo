package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanUp() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanUp()", r)
	}
	defer wg.Done()
}

func say(s string) {
	defer cleanUp()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}

}

func main() {
	wg.Add(1)
	go say("1")
	wg.Add(1)
	go say("2")
	wg.Add(1)
	go say("3")
	wg.Wait()

	time.Sleep(time.Second)
}
