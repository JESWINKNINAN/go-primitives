package main

import "fmt"

func main() {

	myMapName := make(map[int]string)
	myMapName[1] = "jeswin"
	myMapName[2] = "ramzeen"
	myMapName[3] = "sony"

	fmt.Println(myMapName[3])

	for key, value := range myMapName {

		fmt.Println(key, value)
	}

}
