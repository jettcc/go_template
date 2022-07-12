// 字典树模板
package main

import "fmt"

// Trie 树结构
type Trie struct {
	root *trieNode
}

// Trie 树节点
type trieNode struct {
	char     string             // Unicode 字符
	isEnding bool               // 是否是单词结尾
	children map[rune]*trieNode // 该节点的子节点字典
}

// 初始化 Trie 树节点
func NewTrieNode(char string) *trieNode {
	return &trieNode{
		char:     char,
		isEnding: false,
		children: make(map[rune]*trieNode),
	}
}

// 初始化 Trie 树
func NewTrie() *Trie {
	trieNode := NewTrieNode("/")
	return &Trie{trieNode}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			value = NewTrieNode(string(code))
			node.children[code] = value
		}
		node = value
	}
	node.isEnding = true
}

func (t *Trie) Find(word string) bool {
	node := t.root
	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			return false
		}
		node = value
	}
	return node.isEnding
}

func (t *Trie) FindWith(word string) bool {
	node := t.root
	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			return false
		}
		node = value
	}
	return true
}

// 实例
func test() {
	trie := NewTrie()
	words := []string{"Golang", "Language", "Trie", "Go"}
	for _, word := range words {
		trie.Insert(word)
	}
	term := "Golang"
	fmt.Println(trie.Find(term))
}
