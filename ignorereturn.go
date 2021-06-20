package main

import "fmt"

func main() {
	a, _, b := GetSony(3, 8, 4)
	fmt.Println(a, b)

}

func GetSony(a int, b int, c int) (int, int, int) {
	return a, b, c
}
