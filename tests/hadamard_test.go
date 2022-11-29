package tests

import (
	"fmt"
	"testing"

	"github.com/dataStructures/example/models"
)

func TestHadamardProduct(t *testing.T) {
	fmt.Printf("Testing Hadamard Product...\n\n")

	// Example
	matrixA := [][]int{
		{1, 6},
		{2, 3},
	}

	matrixB := [][]int{
		{5, 1},
		{3, 2},
	}

	fmt.Printf("MatrixA: %v \n", matrixA)
	fmt.Printf("MatrixB: %v \n", matrixB)

	var posibleError error
	var resultingMatrix [][]int
	resultingMatrix, posibleError = models.CalculateHadamardProduct(matrixA, matrixB)

	if posibleError != nil {
		panic(posibleError)
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix: %v \n", resultingMatrix)
}
