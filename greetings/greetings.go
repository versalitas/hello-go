package greetings

import (
	"errors"
	"fmt"
	"math/rand"
    "time"
)

func Hello(name string) (string, error) {
	
	//if no name return error
	if name == "" {
		return "", errors.New("Empty name")
	}
	
	//else return randomized greeting message and nil
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//randomformat returns a randomly selected greeting
func randomFormat() string {
	//declare formats, slice of greetings
	formats :=[string]{
		"Hi there, %v. Welcome",
		"Great to see you, %v!",
        "Hail, %v! Well met!",
     }

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}


 //return randomly selected message format
 //randomly generated index based on slice's length
 return formats[rand.Intn(len(formats))]

 } 
 
