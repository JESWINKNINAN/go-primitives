package main

import "fmt"

func main() {
	a := 2
	for i := true; i; i = (a < 3) {
		fmt.Println(a)
	}
}
