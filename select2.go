package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	caneclCh := make(chan struct{})
	dataCh := make(chan int)

	go scanWait(caneclCh)

	go func(cancelCh <-chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh:
				close(dataCh)
				return
			case dataCh <- val:
				val++
				time.Sleep(time.Second)
			}
		}
	}(caneclCh, dataCh)

	for val := range dataCh {
		fmt.Println("REad", val)
	}
}

func scanWait(cancel chan<- struct{}) {
	os.Stdin.Read(make([]byte, 1))
	cancel <- struct{}{}
}
