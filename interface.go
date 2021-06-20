package main

import "fmt"

type values struct {
	a, b int
}
type myInterface interface {
	sum() int
	sub() int
	multi() int
	div() int
}

func main() {
	v := values{1, 3}
	totalOutput(v)
}

func (c values) sum() int {

	return c.a + c.b
}

func (c values) sub() int {
	return c.a - c.b
}
func (c values) multi() int {
	return c.a * c.b
}
func (c values) div() int {
	return c.a / c.b
}

func totalOutput(v myInterface) {
	fmt.Println(v.sum())
	fmt.Println(v.sub())
	fmt.Println(v.multi())
	fmt.Println(v.div())

}
