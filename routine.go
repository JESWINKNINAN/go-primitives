package main

import "fmt"

var myChannel = make(chan int)

func main() {
	a, b := 3, 15
	go sum(a, b)
	abc := <-myChannel //Receiver
	fmt.Println(abc)
}

func sum(a, b int) {

	d := a + b
	myChannel <- d //sender
}
