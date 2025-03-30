package datastructure

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: map[rune]*TrieNode{},
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Search(word string) bool {
	root := t.root
	for _, c := range word {
		if root.children[c] == nil {
			return false
		}
		root = root.children[c]
	}

	return root.isEnd
}

func (t *Trie) Insert(word string) {
	root := t.root
	for _, c := range word {
		if root.children[c] == nil {
			root.children[c] = NewTrieNode()
		}
		root = root.children[c]
	}

	root.isEnd = true
}

func (t *Trie) Delete(word string) {
	t.delete(t.root, word, 0)
}

func (t *Trie) delete(root *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !root.isEnd {
			return false
		}
		root.isEnd = false
		return len(root.children) == 0
	}

	c := rune(word[index])
	if root.children[c] == nil {
		return false
	}

	if t.delete(root.children[c], word, index+1) {
		delete(root.children, c)
		return len(root.children) == 0
	}

	return false
}

// Time complexity: O(n)
// Space complexity: O(n)
