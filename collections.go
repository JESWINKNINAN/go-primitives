package main

import "fmt"

type mine interface {
	sum() int
	substract() int
}

type value struct {
	a int
	b int
}

func main() {
	a := value{3, 1}
	calculator(a)
}
func calculator(cal mine) {
	fmt.Println(cal.sum())
	fmt.Println(cal.substract())
}
func (v value) sum() int {
	return v.a + int(v.b)
}

func (v value) substract() int {
	return v.a - v.b
}
