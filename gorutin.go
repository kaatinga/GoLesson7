package main

import (
	"fmt"
	"runtime"
	"strings"
)

var gortineCount = 5

func main() {
	for i := 0; i < gortineCount; i++ {
		go func(th int) {
			for j := 0; j < 5; j++ {
				fmt.Printf(printMessage(th, j))
				//time.Sleep(time.Microsecond)
				runtime.Gosched()
			}
		}(i)
	}

	fmt.Scanln()
}

func printMessage(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", gortineCount-in),
		"Thread", in,
		"Iteration", j, strings.Repeat("<", j))
}
