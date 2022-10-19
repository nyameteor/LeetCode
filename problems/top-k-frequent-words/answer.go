package main

import (
	"container/heap"
	"fmt"
)

// Min-Heap approach, it's more efficient then Max-Heap in most conditions.
// Time: O(n log(k))
// Space: O(n)

func topKFrequent(words []string, k int) []string {
	wordCount := map[string]int{}
	for _, word := range words {
		wordCount[word]++
	}

	h := &WordHeap{}
	for word, count := range wordCount {
		heap.Push(h, entry{word, count})

		// Let min-heap keep the k maximum items.
		if len(*h) > k {
			heap.Pop(h)
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(entry).word
	}

	return res
}

type entry struct {
	word  string
	count int
}

type WordHeap []entry

func (h WordHeap) Len() int {
	return len(h)
}

func (h WordHeap) Less(i, j int) bool {
	if h[i].count < h[j].count {
		return true
	}
	if h[i].count == h[j].count {
		return h[i].word > h[j].word
	}
	return false
}

func (h WordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WordHeap) Push(x interface{}) {
	*h = append(*h, x.(entry))
}

func (h *WordHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func main() {
	var words []string
	var k int

	words = []string{"i", "love", "leetcode", "i", "love", "coding"}
	k = 2
	fmt.Println(topKFrequent(words, k))

	words = []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	k = 4
	fmt.Println(topKFrequent(words, k))
}
