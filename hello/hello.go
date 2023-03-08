package main

import (
	"fmt"
	"log"

	"github.com/versalitas/hello-go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//declare slice of names
	names := []string{"Sue", "Ivy", "Dan"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
