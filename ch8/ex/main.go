package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	switch {
	case n < 0:
		return -1
	case n == 0, n == 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)

	}

}

func spinner(delay time.Duration) {
	for {
		for _, r := range []byte{'-', '\\', '|', '/'} {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(100 * time.Millisecond)
	x := fib(100)
	fmt.Println("Found Fib:", x)
}
