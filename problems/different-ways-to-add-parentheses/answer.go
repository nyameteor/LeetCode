package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func diffWaysToCompute(expression string) []int {
	if len(expression) == 0 {
		return []int{}
	}

	if len(expression) == 1 ||
		(len(expression) == 2 && unicode.IsDigit(rune(expression[0]))) {
		num, err := strconv.Atoi(expression)
		if err != nil {
			return []int{}
		}
		return []int{num}
	}

	results := []int{}
	for i, r := range expression {
		if unicode.IsDigit(r) {
			continue
		}

		lResults := diffWaysToCompute(expression[0:i])
		rResults := diffWaysToCompute(expression[i+1:])
		op := expression[i]
		for _, x := range lResults {
			for _, y := range rResults {
				var res int
				if op == '+' {
					res = x + y
				} else if op == '-' {
					res = x - y
				} else {
					res = x * y
				}
				results = append(results, res)
			}
		}
	}
	return results
}

func main() {
	var expression string
	var res []int

	expression = "2-1-1"
	res = diffWaysToCompute(expression)
	// [0,2]
	fmt.Println(res)

	expression = "2*3-4*5"
	res = diffWaysToCompute(expression)
	// [-34,-14,-10,-10,10]
	fmt.Println(res)
}
