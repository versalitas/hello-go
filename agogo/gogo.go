package main

import "fmt"

// struct
type Person struct {
	name, surname string
}

// struct literals
var (
	person1 = Person{"Arn", "O"}
	person2 = Person{"Sue", "E"}
)

func main() {
	//array literal
	team := [5]string{"a", "b", "c", "d", "e"}

	//slice literal, same as array literal but without length
	q := []int{2, 3, 5, 7, 11, 13}

	//slcie with make... array, length, capacity
	b := make([]int, 0, 5)

	//slice made from array defines start (including) and end index (excluded)
	var juniors []string = team[1:4]
	fmt.Println(person1, person2)
	fmt.Println(person1.name)
	fmt.Println(team, q, b)
	fmt.Println(juniors)

	for i, v := range q {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}
