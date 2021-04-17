package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()

	return c
}

func receive(q, c <-chan int) {
	switch {
	case <-c:

	}
}
