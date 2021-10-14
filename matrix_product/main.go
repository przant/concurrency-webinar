package main

import (
	"fmt"
	"math/rand"
	"time"
)

type StepOne struct {
	Row    int
	Vector []int64
}

type StepTwo struct {
	Row, Col int
	Result   int64
}

const (
	workers     = 8
	matrixRange = 1024
)

func main() {
	matrixA := make([][]int64, matrixRange)
	matrixB := make([][]int64, matrixRange)
	matrixC := make([][]int64, matrixRange)

	stepOneChan := make(chan StepOne)
	stepTwoChan := make(chan StepTwo)
	matrixFillerChan := make(chan bool)

	start := time.Now()

	go fillMatrix(matrixA, matrixFillerChan)
	go fillMatrix(matrixB, matrixFillerChan)
	go fillMatrix(matrixC, matrixFillerChan)

	for matrixFiller := 1; matrixFiller <= 3; matrixFiller++ {
		<-matrixFillerChan
	}
	close(matrixFillerChan)

	for worker := 1; worker <= workers; worker++ {
		go workerProcess(matrixB, stepOneChan, stepTwoChan)
	}

	go matrixProduct(matrixA, stepOneChan)

	for result := 1; result <= matrixRange*matrixRange; result++ {
		results := <-stepTwoChan
		matrixC[results.Row][results.Col] = results.Result
	}
	close(stepOneChan)
	close(stepTwoChan)

	// fmt.Printf("\n\n")
	// fmt.Printf("MatrixA = %v\n\n", matrixA)
	// fmt.Printf("MatrixB = %v\n\n", matrixB)
	// fmt.Printf("MatrixC = %v\n\n", matrixC)

	duration := time.Since(start)
	duration = duration.Truncate(time.Millisecond)

	fmt.Printf("%s\n", duration)
}

func fillMatrix(matrix [][]int64, c chan bool) {
	for row := 0; row < matrixRange; row++ {
		matrix[row] = make([]int64, matrixRange)
		for col := 0; col < matrixRange; col++ {
			matrix[row][col] = rand.Int63n(9) + 1
		}
	}
	c <- true
}

func matrixProduct(mA [][]int64, stepOneChan chan<- StepOne) {
	for row := 0; row < matrixRange; row++ {
		pipeOneStruct := StepOne{
			Row:    row,
			Vector: mA[row],
		}
		go func(pipeOneData StepOne, stepOne chan<- StepOne) {
			stepOne <- pipeOneData
		}(pipeOneStruct, stepOneChan)
	}
}

func workerProcess(mB [][]int64, stepOneChan <-chan StepOne, stepTwoChan chan<- StepTwo) {
	for stepOne := range stepOneChan {
		for col := 0; col < matrixRange; col++ {
			stepTwo := StepTwo{
				Row: stepOne.Row,
				Col: col,
			}
			for index := 0; index < matrixRange; index++ {
				// fmt.Printf("(%d,%d) = (%d,%d) * (%d,%d) = %d * %d = %d\n\n",
				// 	stepTwo.Row, stepTwo.Col,
				// 	stepTwo.Row, index,
				// 	stepTwo.Col, index,
				// 	stepOne.Vector[index],
				// 	mB[index][col],
				// 	stepOne.Vector[index]*mB[index][col])
				stepTwo.Result += stepOne.Vector[index] * mB[index][col]
			}
			stepTwoChan <- stepTwo
		}
	}
}
