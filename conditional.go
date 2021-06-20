package main

import "fmt"

func main() {

	i := 20

	a, b := 0, 1

	if a+b > 2 {
		fmt.Println(i)
	} else if a+b < 2 {
		fmt.Println(i)
	} else {
		fmt.Println("failed")
	}

}
