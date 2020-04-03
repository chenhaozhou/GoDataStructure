package Trie

type node struct {
	isWord bool
	next   map[byte]*node
}

type Trie struct {
	size int
	root *node
}

func NewTrie() *Trie {
	return &Trie{size: 0, root: &node{isWord: false, next: make(map[byte]*node)}}
}

func (t *Trie) Add(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := cur.next[word[i]]; !ok {
			cur.next[word[i]] = &node{isWord: false, next: make(map[byte]*node)}
		}
		cur = cur.next[word[i]]
	}
	if false == cur.isWord {
		cur.isWord = true
		t.size++
	}
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := cur.next[word[i]]; !ok {
			return false
		}
		cur = cur.next[word[i]]
	}
	return cur.isWord
}

func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		if _, ok := cur.next[prefix[i]]; !ok {
			return false
		}
		cur = cur.next[prefix[i]]
	}
	return true
}

func (t *Trie) Delete(word string) bool {
	if "" == word {
		return false
	}
	res := del(t.root, word)

	if res {
		t.size--
	}

	return res
}

func del(node *node, word string) bool {
	if 0 == len(word) {
		if false == node.isWord {
			return false
		} else {
			node.isWord = false
			return true
		}
	}
	if _, ok := node.next[word[0]]; ok {
		res := del(node.next[word[0]], word[1:])
		child := node.next[word[0]]
		if !child.isWord && len(child.next) == 0 {
			delete(node.next, word[0])
		}
		return res
	} else {
		return false
	}
}

func (t *Trie) GetSize() int {
	return t.size
}
