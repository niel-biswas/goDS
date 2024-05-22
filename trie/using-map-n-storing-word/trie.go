package using_map_n_storing_word

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[byte]*TrieNode
	words    []string
}

// Trie represents the Trie itself
type Trie struct {
	root *TrieNode
}

// NewTrieNode initializes a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

// NewTrie initializes a new Trie
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Insert adds a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
		node.words = append(node.words, word)
	}
}

// FindWordsWithPrefix returns all words in the Trie with a given prefix
func (t *Trie) FindWordsWithPrefix(prefix string) []string {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, exists := node.children[char]; !exists {
			return nil
		}
		node = node.children[char]
	}
	return node.words
}
