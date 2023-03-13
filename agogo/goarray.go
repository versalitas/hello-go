package main

import "fmt"

func main() {
	//have to declare number of elements
	var arr [5]int

	arr2 := [5]int{3, 5, 6}
	sum := 0
	//iteration with for

	for i := 0; i < len(arr2); i++ {
		sum += i
		fmt.Println(sum)
	}

	fmt.Println(arr)
	fmt.Println(sum)

}
