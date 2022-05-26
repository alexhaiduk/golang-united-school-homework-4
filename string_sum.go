package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	var operand1 string = ""
	var operand2 string = ""
	var operand3 string = ""
	var operator string = ""
	var op1, op2 int
	var temp string = ""
	var minus bool = false
	var found bool = false
	if len(input) < 1 {
		//fmt.Errorf("%w", errorEmptyInput)
		return "", errorEmptyInput
	} else {
		for i, v := range input {
			if i == 0 && strings.Contains(string(v), "-") {
				minus = true
			} else {
				temp += string(v)
			}
			if i > 0 && strings.ContainsAny(string(v), "-+") && operator == "" {
				operator = string(v)
			}
		}
		if operator == "" {
			//fmt.Errorf("%w", errorNotTwoOperands)
			return "", errorNotTwoOperands
		} else {
			operand1, operand2, found = strings.Cut(temp, operator)
			if found && strings.ContainsAny(operand2, "-+") {
				operand3, _, found = strings.Cut(operand2, "-")
				if found && operand3 != "" {
					//fmt.Errorf("%w", errorNotTwoOperands)
					return "", errorNotTwoOperands
				}
				operand3, _, found := strings.Cut(operand2, "+")
				if found && operand3 != "" {
					//fmt.Errorf("%w", errorNotTwoOperands)
					return "", errorNotTwoOperands
				}
			}
		}
		if minus {
			operand1 = "-" + operand1
		}
		op1, err = strconv.Atoi(operand1)
		fmt.Println(op1)
		if err != nil {
			//fmt.Errorf("%w", err)
			return "", err
		}
		fmt.Println(operator)
		op2, err = strconv.Atoi(operand2)
		fmt.Println(op2)
		if err != nil {
			//fmt.Errorf("%w", err)
			return "", err
		}
	}
	if operator == "-" {
		return strconv.Itoa(op1 - op2), nil
	} else if operator == "+" {
		return strconv.Itoa(op1 + op2), nil
	}
	return "", nil
}
