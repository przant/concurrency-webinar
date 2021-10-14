package main

import "fmt"

func main() {
	value := 3
	firstChan := make(chan int)
	secondChan := make(chan int)

	go square(value, firstChan)
	go cubic(firstChan, secondChan)

	fmt.Printf("The result is: %d\n", <-secondChan)

	close(firstChan)
	close(secondChan)

}

func square(value int, pipeOne chan<- int) {
	pipeOne <- value
	pipeOne <- value * value
}

func cubic(pipeOne <-chan int, pipeTwo chan<- int) {
	value := <-pipeOne
	squareValue := <-pipeOne
	pipeTwo <- squareValue * value
}
