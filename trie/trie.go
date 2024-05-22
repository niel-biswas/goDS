package trie

const ENGLISH_ALPHABETS = 26

// TrieNode represents a node in the Trie
type TrieNode struct {
	children [ENGLISH_ALPHABETS]*TrieNode
	isEnd    bool
}

// Trie represents the Trie itself
type Trie struct {
	root *TrieNode
}

// NewTrieNode initializes a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

// NewTrie initializes a new Trie
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Insert adds a word into the Trie
func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, rune := range word {
		// convert rune to index
		charAtIndex := rune - 'a'
		if currentNode.children[charAtIndex] == nil {
			currentNode.children[charAtIndex] = &TrieNode{}
		}
		currentNode = currentNode.children[charAtIndex]
	}
	currentNode.isEnd = true
}

// Search returns true/false if the word exists in the Trie
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

// StartsWith returns true/false if any word in the Trie has the given prefix
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

/*func main() {
	myTrie := NewTrie()
	myTrie.Insert("aragorn")
	myTrie.Insert("googly")
	myTrie.Insert("organic")
	fmt.Println(myTrie.Search("googly"))
	fmt.Println(myTrie.StartsWith("goog"))
}*/
