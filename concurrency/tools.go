package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 3)

	go func(input chan string) {
		fmt.Println("sending to channel")
		input <- "hello1"
		fmt.Println("sending to channel")

		input <- "hello2"
		fmt.Println("sending to channel")

		input <- "hello3"

	}(c)
	time.Sleep(5 * time.Second)
	fmt.Println("recieving from channel")

	for greetings := range c {
		fmt.Println("greetings recieved:")
		fmt.Println(greetings)
	}
}
