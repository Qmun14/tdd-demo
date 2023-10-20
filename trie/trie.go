package trie

const AlphabetSize = 26

// Node represents each node in The Trie
type Node struct {
	Children [AlphabetSize]*Node
	IsEnd    bool
}

// a construct from Trie struct
func NewTrie() *Trie {
	return &Trie{Root: &Node{}}
}

// Trie represent a trie and has a pointer to the root node
type Trie struct {
	Root *Node
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) error {
	wordLength := len(w)
	currNode := t.Root
	for i := 0; i < wordLength; i++ {
		charIdx := w[i] - 'a'
		if currNode.Children[charIdx] == nil {
			currNode.Children[charIdx] = &Node{}
		}
		currNode = currNode.Children[charIdx]
	}
	currNode.IsEnd = true
	return nil
}

// Search will take in a word and Return true if that word is included int the Trie
func (t Trie) Search(w string) bool {
	currNode := t.Root
	for i := range w {
		charIdx := w[i] - 'a'
		if currNode.Children[charIdx] == nil {
			return false
		}
		currNode = currNode.Children[charIdx]
	}
	if currNode.IsEnd {
		return true
	}
	return false
}
