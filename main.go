package main

import (
	"fmt"
	"log"

	"github.com/caleb-mwasikira/intro_dsa/algos"
)

func main() {
	infix := "A+2*3"

	postfix, err := algos.InfixToPostFix(infix)
	if err != nil {
		log.Fatal(err)
	}

	expectedResult := 7
	result, err := algos.EvaluateInfixExpression(
		infix,
		map[rune]int{
			'A': 1,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("infix: ", infix)
	fmt.Println("postfix: ", postfix)
	fmt.Println("expected result: ", expectedResult)
	fmt.Println("result: ", result)
}
