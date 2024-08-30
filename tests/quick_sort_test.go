package tests

import (
	"fmt"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/alg"
)

func TestQuickSort(t *testing.T) {
	for index, test := range Tests {
		testName := fmt.Sprintf("testing quickSort: %v", index)

		t.Run(testName, func(t *testing.T) {
			testOutput := test.Input
			alg.QuickSort(testOutput, 0, len(testOutput)-1)

			if !slices.Equal(testOutput, test.ExpectedOutput) {
				t.Errorf("expected %#v but got %#v", test.ExpectedOutput, testOutput)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := Tests[0]
		testOutput := test.Input
		alg.QuickSort(testOutput, 0, len(testOutput)-1)
	}
}
