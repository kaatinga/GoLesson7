package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
LOOP:
	for {
		select {
		case value := <-ch1:
			fmt.Println("ch1 val", value)
		case value2 := <-ch2:
			fmt.Println("ch2 val", value2)
		default:
			break LOOP
		}
	}

}
