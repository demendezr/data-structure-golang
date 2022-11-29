package tests

import (
	"fmt"
	"testing"

	"github.com/dataStructures/example/models"
)

func TestReshape(t *testing.T) {
	fmt.Printf("Testing Reshape...\n\n")
	var tensor models.Tensor
	tensor.Range = 1

	/*
		// Example 1: please quit comments to run this example and comment Example 2
		tensor.Shape = []int{1, 8}
		tensor.Data = [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8},
		}

		var newShape []int = []int{4, 2}
	*/

	// Example 2:
	tensor.Shape = []int{1, 4}
	tensor.Data = [][]int{
		{1, 2, 3, 4},
	}

	var newShape []int = []int{2, 2}

	fmt.Printf("Shape: %v \n", tensor.Shape)
	fmt.Printf("Tensor: %v \n", tensor.Data)

	var posibleError error
	tensor, posibleError = models.Reshape(tensor, newShape)

	if posibleError != nil {
		panic(posibleError)
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Shape: %v \n", tensor.Shape)
	fmt.Printf("New Tensor Reshaped: %v \n", tensor.Data)
}
