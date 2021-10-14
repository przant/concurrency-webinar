package main

import (
	"fmt"
	"runtime"
)

// Task1 adds one to number, which is a pointer to int
func Task1(number *int) {
	*number++ // Operation p1
}

// Task2 adds one to number, which is a pointer to int
func Task2(number *int) {
	*number++ // Operation p2
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	for test := 1; test <= 5; test++ {
		for i := 1; i <= 1000; i++ {
			// The data race occur when within Task1 and Task2 both write to
			// the same share variable, because is not guarantee about which
			// operation goes first: p1, p2, both or none of them.
			//
			// The display values show at certain execution the variable
			// 'shareVariable' increases in one unit

			shareVariable := 1
			go Task2(&shareVariable)
			go Task1(&shareVariable)

			// Posible print values are 1 and 2 for this example
			fmt.Print(shareVariable)
		}
		fmt.Println()
		fmt.Println()
	}
}
