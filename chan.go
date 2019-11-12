package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(in chan int) {
		value := <-in
		fmt.Println("From channel value", value)
	}(ch)

	ch <- 50
	//ch<-100

	fmt.Println("Done send in channel")
	fmt.Scanln()
}
