package main

import (
	"fmt"

	"../greetings"
)

func main() {
	fmt.Println("Hello World!")
	message := greetings.Hello("Sue")
	fmt.Println(message)
}
