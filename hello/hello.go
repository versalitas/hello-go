package main

import (
	"fmt"
	"log"

	"github.com/versalitas/hello-go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Sue")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
