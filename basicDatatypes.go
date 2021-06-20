package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var sampleBool bool
	sampleBool = false
	fmt.Println(sampleBool)
	var sampleString string
	sampleString = "a"
	fmt.Println(sampleString)
	sampleBoolSize := unsafe.Sizeof(sampleBool)
	sampleStringSize := unsafe.Sizeof(sampleString)
	fmt.Println(sampleBoolSize, sampleStringSize)

	var sampleInt uint
	sampleIntSize := unsafe.Sizeof(sampleInt)
	fmt.Println(sampleInt, sampleIntSize)

}
