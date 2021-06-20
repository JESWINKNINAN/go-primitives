package main

import (
	"flag"
	"fmt"

	"calc/pkg/substract"
	"calc/pkg/sum"
)

func main() {

	var a = flag.Int("firstValue", 1, "integer value for first input")
	var b = flag.Int("secondValue", 2, "integer value for second input")
	flag.Parse()
	sumValue := sum.Sumfunc(*a, *b)
	substractValue := substract.Substractfunc(*a, *b)
	fmt.Println(sumValue, substractValue)
}
