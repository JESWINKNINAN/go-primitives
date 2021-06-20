package main

import "fmt"

func main() {

	b := [3]int{1, 2, 4}
	stringArray := [5]string{"abbc", "cab", "dab"}
	var a [4]int           //arrary declaration
	a = [4]int{1, 3, 4, 6} //assigning values
	fmt.Println(a[1])
	fmt.Println(b[2])
	fmt.Printf("%T \n", b)
	fmt.Println("\n", stringArray[2])*/

	var sliceArray []int //Slice
	sliceArray = b[0:2]
	for i := 0; i < len(sliceArray); i++ {
		fmt.Println(sliceArray[i])
	}

	var multiDimSlice [3][3]int

	multiDimSlice = [3][3]int{[3]int{1, 2, 3}, [3]int{4, 6, 7}, [3]int{0, 9, 8}}

	fmt.Println(multiDimSlice[0][2])

}
