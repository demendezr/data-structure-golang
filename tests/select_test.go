package tests

import (
	"fmt"
	"testing"

	"github.com/dataStructures/example/models"
)

func TestIndexSelect(t *testing.T) {
	fmt.Printf("Testing Index Select...\n\n")

	/*
		Example 1:
		=> IndexSelect([1, 2, 3, 4], 0, [0, 0, 2])
		[1, 1, 3]
	*/
	//dimension1 := 0
	//indexes1 := []int{0, 0, 2}

	var tensor1 models.Tensor
	tensor1.Range = 1
	tensor1.Shape = []int{1, 4}
	tensor1.Data = [][]int{
		{1, 2, 3, 4},
	}

	/*
		Example 2:
		=> IndexSelect([[1, 2], [3, 4]], 0, [0])
		[[1, 2]]
	*/

	//dimension2 := 0
	//indexes2 := []int{0}

	var tensor2 models.Tensor
	tensor2.Range = 1
	tensor2.Shape = []int{2, 2}
	tensor2.Data = [][]int{
		{1, 2},
		{3, 4},
	}

	/*
		Example 3:
		=> IndexSelect([[1, 2], [3, 4]], 1, [0, 0, 1, 1])
		[[1, 1, 2, 2], [3, 3, 4, 4]]
	*/

	dimension3 := 1
	indexes3 := []int{0, 0, 1, 1}

	var tensor3 models.Tensor
	tensor3.Range = 1
	tensor3.Shape = []int{2, 2}
	tensor3.Data = [][]int{
		{1, 2},
		{3, 4},
	}

	// Please change the tensor, dimensions and indexes according to the example you want to run
	/* If you change the values, please comment the values unused because golang does not allow you
	to declare a variable and not use it */
	tensor := tensor3
	dimension := dimension3
	indexes := indexes3

	fmt.Printf("Tensor: %v \n", tensor.Data)
	fmt.Printf("Dimension: %v \n", dimension)
	fmt.Printf("Indexes: %v \n", indexes)

	var posibleError error
	var resultingMatrix [][]int
	resultingMatrix, posibleError = models.IndexSelect(tensor, dimension, indexes)

	if posibleError != nil {
		panic(posibleError)
	}

	fmt.Printf("\nResults: \n")
	fmt.Printf("New Matrix with selected indexes: %v \n", resultingMatrix)
}
