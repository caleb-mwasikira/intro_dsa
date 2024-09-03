package tests

import (
	"fmt"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/alg"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

func TestSelectionSort(t *testing.T) {
	for index, test := range Tests {
		testName := fmt.Sprintf("test %v", index)

		t.Run(testName, func(t *testing.T) {
			actualOutput := alg.SelectionSort(test.Input)
			if !slices.Equal(actualOutput, test.ExpectedOutput) {
				t.Errorf("expected %#v, but got %#v\n", test.ExpectedOutput, actualOutput)
			}
		})
		fmt.Println()
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	arr := utils.GenRandomArray(ARRAY_MAX, 0, 100)

	for i := 0; i < b.N; i++ {
		alg.SelectionSort(arr)
	}
}