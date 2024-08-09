package tests

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Test struct {
	Input          []int
	ExpectedOutput []int
}

var (
	ARRAY_MAX int = 100

	Tests []Test = []Test{
		{
			Input:          []int{6, 40, 61, 44, 41, 80, 20, 11, 24, 19},
			ExpectedOutput: []int{6, 11, 19, 20, 24, 40, 41, 44, 61, 80},
		},
		{
			Input:          []int{85, 45, 85, 74, 25, -16, -40, 81, 65, 66},
			ExpectedOutput: []int{-40, -16, 25, 45, 65, 66, 74, 81, 85, 85},
		},
		{
			Input:          []int{-85, -45, -85, -74, -25, -16, -40, -81, -65, -66},
			ExpectedOutput: []int{-85, -85, -81, -74, -66, -65, -45, -40, -25, -16},
		},
	}
)

func init() {
	// get ARRAY_MAX value from env variables or set default to 100
	val := os.Getenv("ARRAY_MAX")
	arrayMax, err := strconv.ParseInt(val, 10, 32)
	if err != nil || arrayMax <= 0 {
		log.Println("missing/invalid env variable ARRAY_MAX")
	} else {
		ARRAY_MAX = int(arrayMax)
	}
	fmt.Printf("ARRAY_MAX=%v\n", ARRAY_MAX)
}
