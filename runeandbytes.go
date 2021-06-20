package main

import (
	"fmt"
)

func main() {

	b := []rune("ഹലോ")

	for index, value := range b {

		fmt.Printf("%v %v \n", index, value)
	}

	fmt.Printf("Unicode translated value: %c \n", b)
	fmt.Printf("Value of unicode value: %v \n", b)
}
