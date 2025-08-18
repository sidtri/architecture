package main

import "fmt"

// Using trie node, design and implement a simple search functionality.
//

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd: false,
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}

	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}

	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}

	return true
}

func Search () {
	words := []string{"the", "a", "answer", "any", "by", "bye", "their", "abc"}

	trie := NewTrie()

	for _, w := range words {
		trie.Insert(w)
	}

	fmt.Println("Searching for some strings using Trie")

  fmt.Println(trie.Search("there"))      // false
	fmt.Println(trie.StartsWith("by"))     // true
	fmt.Println(trie.Search("the"))        // true
	fmt.Println(trie.Search("thes"))       // false
	fmt.Println(trie.StartsWith("th"))     // true
}
