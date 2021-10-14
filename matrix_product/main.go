package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	matrixRange = 1024
)

func main() {
	matrixA := make([][]int64, matrixRange)
	matrixB := make([][]int64, matrixRange)

	start := time.Now()
	fillMatrix(matrixA)
	fillMatrix(matrixB)

	matrxiProduct(matrixA, matrixB)

	duration := time.Since(start)
	duration = duration.Truncate(time.Millisecond)

	fmt.Printf("%s\n", duration)
}

func fillMatrix(matrix [][]int64) {
	rand.Seed(time.Now().Unix())

	for row := 0; row < matrixRange; row++ {
		matrix[row] = make([]int64, matrixRange)
		for col := 0; col < matrixRange; col++ {
			matrix[row][col] = rand.Int63()%9 + 1
		}
	}
}

func matrxiProduct(mA, mB [][]int64) {
	result := make([][]int64, matrixRange)
	for row := 0; row < matrixRange; row++ {
		result[row] = make([]int64, matrixRange)
		for col := 0; col < matrixRange; col++ {
			for productRC := 0; productRC < matrixRange; productRC++ {
				result[row][col] += mA[row][productRC] * mB[productRC][col]
			}

		}
	}
}
