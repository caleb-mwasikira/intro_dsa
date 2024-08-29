package algos

import (
	"fmt"
	"math"
	"math/big"
	"slices"
	"strconv"
	"strings"

	"github.com/caleb-mwasikira/intro_dsa/ds"
)

var (
	// mathematical operators defined in order of increasing precedence;
	// lowest precedence -> highest precendence
	operatorPrecedence = map[rune]int{
		'^': 10,
		'*': 9,
		'/': 8,
		'-': 7,
		'+': 6,
	}
	operators = []rune{'-', '+', '/', '*', '^', ')', '('}
)

type ErrInvalidExpr struct {
	msg string
}

func (e ErrInvalidExpr) Error() string {
	if len(e.msg) != 0 {
		return fmt.Sprintln("invalid expr: ", e.msg)
	}
	return "invalid infix/prefix/postfix expression"
}

/*
function that converts an infix expression in the form (a char b)
into a postfix expression (ab char) where a, b are operands and
char is the operator.
*/
func InfixToPostFix(infix string) (string, error) {
	postfix := []rune{}

	// step 1. create an empty stack
	stack, err := ds.NewStack[rune](len(infix))
	if err != nil {
		return "", err
	}

	precedenceOf := func(val rune) int {
		return operatorPrecedence[val]
	}

	// step 2. scan the infix expr from left to right
	for _, char := range infix {
		strChar := fmt.Sprintf("%c", char)
		if strChar == " " {
			continue
		}

		isOperand := !slices.Contains(operators, char)
		if isOperand {
			// also appending whitespace to output string; essential
			// in differentiating between 2 operands close to each other
			// e.g AB should be A B
			postfix = append(postfix, char, ' ')
			continue
		}

		// if stack is empty \\ top element is a left bracket \\
		// operator is a left bracket, push operator onto stack
		top, _ := stack.Peek()
		if stack.IsEmpty() || top == '(' || char == '(' {
			_, err = stack.Push(char)
			if err != nil {
				return "", err
			}
			continue
		}

		// if operator is ) bracket, pop operators from the stack
		// and append them to the output string until you encounter
		// a ( bracket. Then pop the ( bracket from the stack
		if char == ')' {
			for {
				top, err = stack.Pop()
				if err != nil {
					break
				}

				if top == '(' {
					break
				}

				postfix = append(postfix, top)
			}
			continue
		}

		// operator has higher precedence than top element
		// on the stack
		if precedenceOf(char) > precedenceOf(top) {
			_, err = stack.Push(char)
			if err != nil {
				return "", err
			}
			continue
		}

		// if the operator has lower or equal precedence than
		// the top element on the stack (condition), pop operators and append
		// them to the output string until condition changes
		if precedenceOf(char) <= precedenceOf(top) {
			for {
				top, err = stack.Pop()
				if err != nil {
					break
				}

				postfix = append(postfix, top)

				var err error
				top, err = stack.Peek()
				if err != nil {
					break
				}

				if precedenceOf(char) > precedenceOf(top) || top == '(' {
					break
				}
			}

			_, err = stack.Push(char)
			if err != nil {
				return "", err
			}
			continue
		}
	}

	// get items left on stack and append them to output string
	// following the correct order of precedence
	items := stack.GetItems()
	postfix = append(postfix, items...)

	return string(postfix), nil
}

func EvaluateInfixExpression(infix string, fillInValues map[rune]int) (float64, error) {
	infix = strings.TrimSpace(infix)
	var bigResult big.Float

	postfix, err := InfixToPostFix(infix)
	if err != nil {
		return 0, err
	}

	stack, err := ds.NewStack[float64](len(infix))
	if err != nil {
		return 0, err
	}

	toDigit := func(char rune) (float64, bool) {
		digit, err := strconv.ParseFloat(fmt.Sprintf("%c", char), 64)
		return digit, err == nil
	}

	// scan the postfix expression from left to right.
	// if character is a an operand (in this case a digit) we push it
	// onto the stack. if character is an operator, we pop two
	// operands from the stack and carry out the operation defined
	// by our operator
	for _, char := range postfix {
		strChar := fmt.Sprintf("%c", char)
		if strChar == " " {
			continue
		}

		isOperator := slices.Contains(operators, char)
		digit, isDigit := toDigit(char)

		if !isDigit && !isOperator {
			// try replacing char with a fill-in-value
			newValue, ok := fillInValues[char]
			if !ok {
				// wtf did you give me???
				return 0, ErrInvalidExpr{
					msg: fmt.Sprintf("unexpected character '%c' found in infix expression %s", char, infix),
				}
			}

			isDigit = true
			digit = float64(newValue)
		}

		if isDigit {
			_, err = stack.Push(float64(digit))
			if err != nil {
				return 0, err
			}
		} else {
			// if we fail to pop exactly two operands from our stack,
			// then the infix expression provided was incorrect
			one, err := stack.Pop()
			if err != nil {
				return 0, ErrInvalidExpr{
					msg: "expected to pop 2 operands but stack was empty",
				}
			}
			two, err := stack.Pop()
			if err != nil {
				return 0, ErrInvalidExpr{
					msg: "expected to pop 2 operands but only got 1",
				}
			}

			// some operations require order eg 3^4 is different from 4^3
			// the last operand to be popped is the base or left hand operand
			a := big.NewFloat(two)
			b := big.NewFloat(one)

			switch char {
			case '^':
				res := math.Pow(two, one)
				bigResult = *big.NewFloat(res)
			case '*':
				bigResult = *new(big.Float).Mul(a, b)
			case '/':
				bigResult = *new(big.Float).Quo(a, b)
			case '-':
				bigResult = *new(big.Float).Sub(a, b)
			case '+':
				bigResult = *new(big.Float).Add(a, b)
			default:
				return 0, ErrInvalidExpr{
					msg: fmt.Sprintf("unexpected character '%c' found in infix expression %s", char, infix),
				}
			}

			// push result of the operation onto the stack
			float64Res, _ := bigResult.Float64()
			_, err = stack.Push(float64Res)
			if err != nil {
				return 0, err
			}
		}
	}

	result, _ := bigResult.Float64()
	return result, nil
}
