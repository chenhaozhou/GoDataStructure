package Trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Add("pandas")
	if !trie.Contains("pandas") {
		t.Error("trie error")
	}
	if trie.Contains("pan") {
		t.Error("trie error:pan not exists")
	}
	trie.Add("pan")
	if !trie.Contains("pan") {
		t.Error("trie error:pan is exists")
	}

	if !trie.IsPrefix("pa") {
		t.Error("is prefix error")
	}
}

func TestTrie_GetSize(t *testing.T) {
	trie := NewTrie()
	if 0 != trie.GetSize() {
		t.Error("getSize error")
	}
	trie.Add("apple")
	trie.Add("pandas")
	if 2 != trie.GetSize() {
		t.Error("getSize error")
	}
}
