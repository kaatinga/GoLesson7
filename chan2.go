package main

import "fmt"

func main() {
	in := make(chan int, 2)

	go func(out chan<- int) {
		for i := 0; i < 10; i++ {
			fmt.Println("Read channel", i)
			out <- i
			fmt.Println("Read completed", i)
		}
		close(out)
	}(in)

	for i := range in {
		fmt.Println("\t val", i)
	}
}
