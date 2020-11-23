package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	go func() {
		messages <- "Hello, World!\n"
	}()

	msg := <-messages
	fmt.Println(msg)
	go func() {
		messages <- "New Information!"
	}()

	newMsg := <-messages
	fmt.Println(newMsg)

}
