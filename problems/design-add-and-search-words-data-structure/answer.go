package main

import (
	"fmt"
	"slices"
)

type Node struct {
	char     byte
	isEnd    bool
	children []*Node
}

type Trie struct {
	roots []*Node
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Add(word string) {
	add(&t.roots, word)
}

func (t *Trie) Search(word string) bool {
	return search(&t.roots, word)
}

func add(nodes *[]*Node, word string) {
	if len(word) == 0 {
		return
	}

	char := word[0]
	isEnd := len(word) == 1

	for _, node := range *nodes {
		if char == node.char {
			if isEnd {
				node.isEnd = true
			}
			add(&node.children, word[1:])
			return
		}
	}

	newNode := &Node{
		char:  char,
		isEnd: isEnd,
	}
	*nodes = append(*nodes, newNode)
	add(&newNode.children, word[1:])
}

func search(nodes *[]*Node, word string) bool {
	if len(word) == 0 {
		return false
	}

	char := word[0]
	isEnd := len(word) == 1

	if char == '.' {
		for _, node := range *nodes {
			if isEnd && node.isEnd {
				return true
			}
			if isEnd && !node.isEnd {
				continue
			}
			if search(&node.children, word[1:]) {
				return true
			}
		}
		return false
	}

	for _, node := range *nodes {
		if char == node.char {
			if isEnd && node.isEnd {
				return true
			}
			if isEnd && !node.isEnd {
				return false
			}
			return search(&node.children, word[1:])
		}
	}

	return false
}

type WordDictionary struct {
	trie *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{
		trie: NewTrie(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.trie.Add(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word)
}

func main() {
	type Input struct {
		operations []string
		parameters [][]string
	}

	testCases := []struct {
		input *Input
		want  []bool
	}{
		{
			input: &Input{
				[]string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
				[][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}},
			},
			want: []bool{false, true, true, true},
		},
		{
			input: &Input{
				[]string{"WordDictionary", "addWord", "addWord", "search", "search", "search", "search", "search", "search"},
				[][]string{{}, {"a"}, {"a"}, {"."}, {"a"}, {"aa"}, {"a"}, {".a"}, {"a."}},
			},
			want: []bool{true, true, false, true, false, false},
		},
		{
			input: &Input{
				[]string{"WordDictionary", "addWord", "addWord", "addWord", "addWord", "search", "search", "addWord", "search", "search", "search", "search", "search", "search"},
				[][]string{{}, {"at"}, {"and"}, {"an"}, {"add"}, {"a"}, {".at"}, {"bat"}, {".at"}, {"an."}, {"a.d."}, {"b."}, {"a.d"}, {"."}},
			},
			want: []bool{false, false, true, true, false, false, true, false},
		},
	}

	for i, tc := range testCases {
		var wd WordDictionary
		got := []bool{}
		for i := range tc.input.operations {
			op := tc.input.operations[i]
			params := tc.input.parameters[i]

			switch op {
			case "WordDictionary":
				wd = Constructor()
			case "addWord":
				wd.AddWord(params[0])
			default:
				res := wd.Search(params[0])
				got = append(got, res)
			}
		}

		if !slices.Equal(got, tc.want) {
			fmt.Printf("Test case %d failed\n", i+1)
			fmt.Printf("Input:\t%v\nGot:\t%v\nWant:\t%v\n\n", tc.input, got, tc.want)
		}
	}
}
