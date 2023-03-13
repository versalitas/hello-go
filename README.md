# hello-go

When working locally

````
go mod edit -replace <git-repository-path>=<local-path>

````
array

var firstArray [5] int

multidimensional

var multiArray [3][3] int

Go arrays don't have substrings, work with slices
The default is **zero** for the low bound and the **length of the slice** for the high bound.

For the array
````
var a [10]int

````
these slice expressions are equivalent:
 
````
a[0:10]
a[:10]
a[0:]
a[:]

````

zero value of slice is **nil**

var emptySlice []int // prints nill

numbers := make([]int, 5 , 10)  (<type>, <length>, <capacity>)

increase slice capacity

append(numbers, 1, 2, 3, 4)

or with copy

//create new slice with bigger capacity
numbers2 := make([]int, 15)
//copy original slice to new
copy(numbers2, number)  // (<target> <original>)




package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))   //prints nil, 0, 0 
	if s == nil {
		fmt.Println("nil!")
	}
}

slices of slices 

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}


MAPS : maps keys to values

var firstMap map[string]int  //map[<type for key>]<type of value>

//adding key/ values

firstMap['age'] = 12
fmt.Println(firstMap['age']) //prints 12


ITERATIONS

i:= 0
sum := 0
for i <= 10 {
	sum += i
	i++
}

fmt.Println(sum)

for i := 0; i < 10 ; i++ {
	sum += i
}

POINTERS

var pointer1 *int

var num := 12
pointer1 = &num

Pointers are usually preferred while passing a struct as an argument or while declaring a method for a defined type.

STRUCTS

collection of different fields

type person struct {
	name string
	age int
}

p1 = person{name: "Sue", age: 5}
person{"Bob", 42}
pp = &person2{name: "Sue", age: 5}