package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	Root *Node
}

type Node struct {
	Children  []*Node
	IsTeminal bool
}

const AlphabetSize = 26

func Constructor() Trie {
	return Trie{
		Root: &Node{
			Children:  make([]*Node, AlphabetSize),
			IsTeminal: false,
		},
	}
}

func (this *Trie) Insert(word string) {
	var insert func(root *Node, word string, i int)
	insert = func(root *Node, word string, i int) {
		if i == len(word) {
			root.IsTeminal = true
			return
		}
		if root.Children[word[i]-'a'] == nil {
			root.Children[word[i]-'a'] = &Node{
				Children:  make([]*Node, AlphabetSize),
				IsTeminal: false,
			}
		}
		insert(root.Children[word[i]-'a'], word, i+1)
	}

	insert(this.Root, word, 0)
}

func (this *Trie) Search(word string) bool {
	var dfs func(root *Node, word string, i int) bool
	dfs = func(root *Node, word string, i int) bool {
		if i >= len(word) {
			return root.IsTeminal
		}
		nextRoot := root.Children[word[i]-'a']
		if nextRoot != nil {
			return dfs(nextRoot, word, i+1)
		}
		return false
	}
	return dfs(this.Root, word, 0)
}

func (this *Trie) StartsWith(prefix string) bool {
	var dfs func(root *Node, prefix string, i int) bool
	dfs = func(root *Node, prefix string, i int) bool {
		if i >= len(prefix) {
			return true
		}
		nextRoot := root.Children[prefix[i]-'a']
		if nextRoot != nil {
			return dfs(nextRoot, prefix, i+1)
		}
		return false
	}
	return dfs(this.Root, prefix, 0)
}

func (this *Trie) String() string {

	var buildString func(root *Node, val rune, builder *strings.Builder, depth int)
	buildString = func(root *Node, val rune, builder *strings.Builder, depth int) {
		if root == nil {
			return
		}

		builder.WriteString(strings.Repeat(" ", depth*2))
		builder.WriteString(fmt.Sprintf("Node{Val: %c, IsTerminal: %v}\n", val, root.IsTeminal))

		for i := 0; i < len(root.Children); i++ {
			buildString(root.Children[i], rune('a'+i), builder, depth+1)
		}
	}

	builder := &strings.Builder{}
	buildString(this.Root, '0', builder, 0)
	return builder.String()
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	var trie Trie
	trie = Constructor()

	trie.Insert("apple")
	// true
	fmt.Println(trie.Search("apple"))
	// false
	fmt.Println(trie.Search("app"))
	// true
	fmt.Println(trie.StartsWith("app"))

	trie.Insert("app")
	// true
	fmt.Println(trie.Search("app"))
}
