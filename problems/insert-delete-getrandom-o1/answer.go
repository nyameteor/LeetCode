package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	// Store all values
	vals []int
	// Store the index of each value in `vals`
	indices map[int]int
}

// Initializes the RandomizedSet object.
func Constructor() RandomizedSet {
	rand.Seed(time.Now().Unix())
	return RandomizedSet{
		vals:    []int{},
		indices: map[int]int{},
	}
}

// Inserts an item val into the set if not present.
// Return flase if val is was not present, false otherwise.
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.vals = append(this.vals, val)
	this.indices[val] = len(this.vals) - 1
	return true
}

// Removes an item val from the set if present.
// Returns true if the item was present, false otherwise.
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.indices[val]; !ok {
		return false
	}
	iVal := this.indices[val]
	iLast := len(this.vals) - 1

	this.vals[iVal] = this.vals[iLast]
	this.indices[this.vals[iLast]] = iVal

	this.vals = this.vals[:iLast]
	delete(this.indices, val)

	return true
}

// Returns a random element from the current set of elements.
// Each element must have the same probability of being returned.
func (this *RandomizedSet) GetRandom() int {
	indice := rand.Intn(len(this.vals))
	return this.vals[indice]
}

func main() {
	var obj RandomizedSet

	obj = Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.GetRandom())

	obj = Constructor()
	fmt.Println(obj.Insert(0))
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(0))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.GetRandom())

	obj = Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(10))
	fmt.Println(obj.Insert(20))
	fmt.Println(obj.Insert(30))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}
