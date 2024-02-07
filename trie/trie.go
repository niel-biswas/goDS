package main

import "fmt"

const ENGLISH_ALPHABETS = 26

type Node struct {
	children [ENGLISH_ALPHABETS]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, rune := range word {
		// convert rune to index
		charAtIndex := rune - 'a'
		if currentNode.children[charAtIndex] == nil {
			currentNode.children[charAtIndex] = &Node{}
		}
		currentNode = currentNode.children[charAtIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, rune := range word {
		// convert rune to index
		charAtIndex := rune - 'a'
		if currentNode.children[charAtIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charAtIndex]
	}
	return currentNode.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t.root
	for _, rune := range prefix {
		// convert rune to index
		charAtIndex := rune - 'a'
		if currentNode.children[charAtIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charAtIndex]
	}
	return true
}

func main() {
	myTrie := NewTrie()
	myTrie.Insert("aragorn")
	myTrie.Insert("googly")
	myTrie.Insert("organic")
	fmt.Println(myTrie.Search("googly"))
	fmt.Println(myTrie.StartsWith("goog"))
}
