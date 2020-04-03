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

func TestTrie_Delete(t *testing.T) {
	trie := NewTrie()
	trie.Add("pandas")
	trie.Add("pan")
	trie.Add("pa")

	if !trie.Contains("pandas") {
		t.Error("pandas is not exists")
	}
	trie.Delete("pandas")
	if trie.Contains("pandas") {
		t.Error("delete pandas error")
	}
	if !trie.Contains("pan") {
		t.Error("error:pan not exists")
	}
	if !trie.Contains("pa") {
		t.Error("error:pa not exists")
	}

	trie.Delete("pa")
	if trie.Contains("pa") {
		t.Error("error: pa exists")
	}

	if !trie.Contains("pan") {
		t.Error("error: pan not exists")
	}

	trie.Add("pa")
	if !trie.Contains("pa") {
		t.Error("error:pa not exists")
	}
	if !trie.Delete("pan") {
		t.Error("error:delete pan filed")
	}
	if trie.Contains("pan") {
		t.Error("error:pan exists")
	}
	if !trie.Contains("pa") {
		t.Error("error:pa not exists")
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
