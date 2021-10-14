package main

import (
	"fmt"
	"sync"
	"time"
)

func ProcessA(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Printf("Hello ")
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("World!\n")
	}
}

func ProcessB(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Printf("Hola, ")
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("Mundo!\n")
	}
}

func main() {
	waitGroup := new(sync.WaitGroup)

	waitGroup.Add(2)

	go ProcessA(waitGroup)
	go ProcessB(waitGroup)

	waitGroup.Wait()
}
