package main

import (
	"fmt"

	//"../greetings"
	"github.com/versalitas/hello-go/greetings"
)

func main() {
	fmt.Println("Hello World!")
	message := greetings.Hello("Sue")
	fmt.Println(message)
}
