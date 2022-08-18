package main

import (
	"fmt"
)

type Symbol struct {
	Sym string
	Val int
}

var (
	symbols = []Symbol{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
)

func intToRoman(num int) string {
	res := ""
	for i := 0; i < len(symbols); {
		symbol := symbols[i]
		if num >= symbol.Val {
			num -= symbol.Val
			res += symbol.Sym
		} else {
			i++
		}
	}

	return res
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
