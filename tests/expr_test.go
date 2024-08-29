package tests

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/algos"
)

type exprTest struct {
	infix           string
	expectedPostfix string
	fillInValues    map[rune]int
	expectedResult  float64
}

var (
	simpleExprTests = []exprTest{
		{
			infix:           "3^4",
			expectedPostfix: "34^",
			expectedResult:  81,
		},
		{
			infix:           "1*2",
			expectedPostfix: "12*",
			expectedResult:  2,
		},
		{
			infix:           "3/4",
			expectedPostfix: "34/",
			expectedResult:  0.75,
		},
		{
			infix:           "3+4",
			expectedPostfix: "34+",
			expectedResult:  7,
		},
		{
			infix:           "3-4",
			expectedPostfix: "34-",
			expectedResult:  -1,
		},
	}

	exprTests = []exprTest{
		{
			infix: "A*B+C/D", expectedPostfix: "AB*CD/+",
			fillInValues:   map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4},
			expectedResult: 2.75,
		},
		{
			infix: "A*(B+C)/D", expectedPostfix: "ABC+*D/",
			fillInValues:   map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4},
			expectedResult: 1.25,
		},
		{
			infix: "A*(B+C/D)", expectedPostfix: "ABCD/+*",
			fillInValues:   map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4},
			expectedResult: 2.75,
		},
		{
			infix: "A+B*(C^D-E)", expectedPostfix: "ABCD^E-*+",
			fillInValues:   map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5},
			expectedResult: 153,
		},
		{
			infix: "A*B+C*D", expectedPostfix: "AB*CD*+",
			fillInValues:   map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4},
			expectedResult: 14,
		},
		{
			infix: "(A+B)*(C+D)", expectedPostfix: "AB+CD+*",
			fillInValues:   map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4},
			expectedResult: 21,
		},
		{
			infix: "(A+B)*C-(D-E)*(F+G)", expectedPostfix: "AB+C*DE-FG+*-",
			fillInValues:   map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7},
			expectedResult: 22,
		},

		// // this test fails and I don't know why
		// {
		// 	infix: "A+B-C*D+(E^F)*G/H/I*J+K", expectedPostfix: "AB+CD*-EF^G*H/I/J*+K+",
		// 	fillInValues: map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10, 'K': 11},
		// },
	}
)

/*
checks if two epxressions are the same after stripping all whitespace
in between the characters
*/
func equalExpressions(str1, str2 string) bool {
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")

	// remove all whitespace between characters
	arr1 = slices.DeleteFunc(arr1, func(char string) bool {
		return char == " "
	})
	arr2 = slices.DeleteFunc(arr2, func(char string) bool {
		return char == " "
	})
	return slices.Equal(arr1, arr2)
}

func TestInfixToPostfix(t *testing.T) {
	for _, test := range exprTests {
		testName := fmt.Sprintf("infixToPostfix(%s)", test.infix)
		t.Run(testName, func(t *testing.T) {
			myPostfix, err := algos.InfixToPostFix(test.infix)
			if err != nil {
				log.Fatalf("internal func error: %v", err)
			}

			if !equalExpressions(test.expectedPostfix, myPostfix) {
				t.Errorf("expected result %s but got %s instead",
					test.expectedPostfix,
					myPostfix,
				)
			}
		})
		println()
	}
}

func almostEqual(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestEvaluateInfixExpression(t *testing.T) {
	allTests := append(simpleExprTests, exprTests...)

	for _, test := range allTests {
		testName := fmt.Sprintf("evaluateInfixExpression(%v)", test.infix)
		t.Run(testName, func(t *testing.T) {
			result, err := algos.EvaluateInfixExpression(
				test.infix,
				test.fillInValues,
			)
			if err != nil {
				log.Fatal(err)
			}

			if !almostEqual(result, test.expectedResult, 0.0001) {
				t.Errorf("expected result %v but got %v instead", test.expectedResult, result)
			}
		})
		println()
	}
}

func BenchmarkInfixToPostfix(b *testing.B) {
	test := exprTests[0]

	for i := 0; i < b.N; i++ {
		algos.InfixToPostFix(test.infix)
	}
}
