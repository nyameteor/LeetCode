package main

import (
	"fmt"
)

type CombinationIterator struct {
	combinations []string
	pos          int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combinations := make([]string, 0)
	subset := make([]byte, 0, combinationLength)

	var generate func(start int)
	generate = func(start int) {
		if len(subset) == combinationLength {
			combinations = append(combinations, string(subset))
			return
		}

		for i := start; i < len(characters); i++ {
			subset = append(subset, characters[i])
			generate(i + 1)
			subset = subset[:len(subset)-1]
		}
	}

	generate(0)

	return CombinationIterator{
		combinations: combinations,
		pos:          0,
	}
}

func (this *CombinationIterator) Next() string {
	res := this.combinations[this.pos]
	this.pos++
	return res
}

func (this *CombinationIterator) HasNext() bool {
	return this.pos < len(this.combinations)
}

func main() {
	itr := Constructor("abc", 2)
	fmt.Println(itr.Next())    // return "ab"
	fmt.Println(itr.HasNext()) // return True
	fmt.Println(itr.Next())    // return "ac"
	fmt.Println(itr.HasNext()) // return True
	fmt.Println(itr.Next())    // return "bc"
	fmt.Println(itr.HasNext()) // return False
}
