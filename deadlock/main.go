package main

import "fmt"

func ProcessA(cA, cB chan int, c chan bool) {
	packet := <-cB
	fmt.Printf("The received value is %d\n", packet)
	cA <- 100
	c <- true
}

func processB(cA, cB chan int, c chan bool) {
	packet := <-cA
	fmt.Printf("The received value is %d\n", packet)
	cB <- 100
	c <- true
}

func main() {
	channelA := make(chan int)
	channelB := make(chan int)
	doneChannel := make(chan bool)

	go ProcessA(channelA, channelB, doneChannel)
	go processB(channelA, channelB, doneChannel)

	for i := 0; i < 2; i++ {
		<-doneChannel
	}
}
