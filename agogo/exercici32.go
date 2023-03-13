package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	nom := scanner()
	fmt.Println(nom)
}

func scanner() string {
	reader := bufio.NewScanner(os.Stdin)
	text, _ := reader.ReadString('\n')
	reader.Scan()
	return reader.Text()

}
