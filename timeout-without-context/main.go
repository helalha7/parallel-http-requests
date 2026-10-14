package main

import (
	"fmt"
	"time"
)

func count(interupt <-chan struct{}) {
	i := 0
	for i <= 100000000000 {
		select {
		case <-interupt:
			return
		default:
			fmt.Println(i)
			i++
		}
	}
}

func countWithTimeout(timeout time.Duration) {
	interupt := make(chan struct{})

	go func() {
		time.Sleep(timeout)
		close(interupt)
	}()
	go count(interupt)

	<-interupt
	return
}

func main() {
	countWithTimeout(2 * time.Second)
}
