package tests

import (
	"fmt"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/algos"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

func TestBubbleSort(t *testing.T) {
	for index, test := range Tests {
		testName := fmt.Sprintf("test %v", index)

		t.Run(testName, func(t *testing.T) {
			actualOutput := algos.BubbleSort(test.Input)
			if slices.Compare(actualOutput, test.ExpectedOutput) != 0 {
				t.Errorf("expected %#v, but got %#v\n", test.ExpectedOutput, actualOutput)
			}
		})
		fmt.Println()
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := utils.GenRandomArray(ARRAY_MAX, 0, 100)

	for i := 0; i < b.N; i++ {
		algos.BubbleSort(arr)
	}
}
