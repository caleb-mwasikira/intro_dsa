package tests

type Test struct {
	Input          []int
	ExpectedOutput []int
}

var (
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

const (
	ARRAY_MAX int = 100
)
