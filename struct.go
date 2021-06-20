package main

import "fmt"

func main() {

	type employee struct {
		eName           string
		age             int
		currentEmployee bool
	}

	a := employee{"sony", 23, true}

	var v *employee
	v = &a
	fmt.Println(v.age)

}
