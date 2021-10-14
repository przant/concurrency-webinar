package main

import (
	"fmt"
	"time"
)

func main() {

	go spinner()

	fmt.Printf("\r\tFibonacci(%d) = %d\n\n", 45, fibonacci(45))
}

func fibonacci(nthNumber int) int {
	if nthNumber < 2 {
		return nthNumber
	}
	return fibonacci(nthNumber-1) + fibonacci(nthNumber-2)
}

func spinner() {
	fmt.Printf("\n")
	for {
		for _, char := range `-\|/` {
			fmt.Printf("\r\t%c", char)
			time.Sleep(250 * time.Millisecond)
		}
	}
}
