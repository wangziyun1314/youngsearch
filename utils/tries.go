package utils

import (
	"fmt"
	"strings"
)

// Node 字典树的节点
type Node struct {
	Term     string
	Children map[string]*Node
	end      bool
	Sentence string
}

// NewNode  新建一个字典树的节点
func NewNode(s string) *Node {
	node := &Node{Term: s}
	m := make(map[string]*Node, 0)
	node.Children = m
	return node
}

// Trie 字典树
type Trie struct {
	Root *Node
}

// NewTries 新建一个字典树
func NewTries() *Trie {
	root := NewNode("\000")
	return &Trie{
		Root: root,
	}
}

// Insert 插入一个单词
func (t *Trie) Insert(word string) error {
	current := t.Root
	newWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for _, s := range newWord {
		term := fmt.Sprintf("%c", s)
		_, ok := current.Children[term]
		if !ok {
			current.Children[term] = NewNode(term)
		}
		current = current.Children[term]
	}
	current.end = true
	current.Sentence = word
	return nil
}

// SearchWord 查找字典树中有没有这个单词
func (t *Trie) SearchWord(word string) bool {
	current := t.Root
	newWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for _, s := range newWord {
		term := fmt.Sprintf("%c", s)
		if current == nil || current.Children[term] == nil {
			return false
		}
		current = current.Children[term]
	}
	return true
}

// FindLastNode 找到这个单词的最后节点，推荐的时候需要用到
func (t *Trie) FindLastNode(word string) *Node {
	current := t.Root
	newWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for _, s := range newWord {
		term := fmt.Sprintf("%c", s)
		if current.Children[term] == nil {
			return current
		}
		current = current.Children[term]
	}
	return current
}

// Recommend 根据输入的单词进行推荐，根据层序遍历进行推荐，默认最多推荐十个
func (t *Trie) Recommend(word string) []string {
	node := t.FindLastNode(word)
	res := make([]string, 0)
	queue := NewQueue()
	queue.Offer(node)
	index := 0
	for !queue.IsEmpty() {
		poll := queue.Poll()
		n := poll.(*Node)
		if n.end {
			res = append(res, n.Sentence)
			index++
			if index == MaxTriesRecommend {
				break
			}
		}
		for _, n2 := range n.Children {
			if n2 != nil {
				queue.Offer(n2)
			}
		}
	}
	return res
}
