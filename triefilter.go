package sx

import (
	"bytes"
	"strings"
	"sync"
	"unicode"
)

type trieNode struct {
	index    int  // 该节点对应的敏感词在词表中的位置
	end      bool // 用于标记是否为一个敏感词结束
	children map[rune]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
	}
}

func (this *trieNode) getNode(r rune) *trieNode {
	var node = this.children[r]
	return node
}

type TrieFilter struct {
	pool     *sync.Pool
	root     *trieNode
	words    []string
	excludes map[rune]struct{}
}

func NewTrieFilter(stock WordStock) *TrieFilter {
	var t = &TrieFilter{}
	t.prepare(stock)
	return t
}

func (this *TrieFilter) prepare(stock WordStock) {
	this.pool = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
	this.root = newTrieNode()
	this.root.index = -1

	var words = stock.ReadAll()

	for index, word := range words {
		if len(word) == 0 {
			continue
		}
		this.addNode(index, word)
	}
	this.words = words
	this.excludes = make(map[rune]struct{})
	return
}

func (this *TrieFilter) addNode(index int, word string) {
	var node = this.root
	var wChars = []rune(strings.TrimSpace(word))

	for _, r := range wChars {
		if unicode.IsSpace(r) {
			continue
		}

		if _, ok := node.children[r]; !ok {
			node.children[r] = newTrieNode()
		}
		node = node.children[r]
	}
	node.index = index
	node.end = true
}

func (this *TrieFilter) skip(r rune) bool {
	// 太影响效率
	if /* unicode.IsSpace(r) || unicode.IsPunct(r) || */ this.inExclude(r) {
		return true
	}
	return false
}

func (this *TrieFilter) inExclude(r rune) bool {
	_, ok := this.excludes[r]
	return ok
}

func (this *TrieFilter) Excludes(items ...rune) {
	for _, item := range items {
		this.excludes[item] = struct{}{}
	}
}

func (this *TrieFilter) Contains(text string) bool {
	var node *trieNode
	var tChars = []rune(text)

	for _, r := range tChars {
		if this.skip(r) {
			continue
		}

		if node != nil {
			node = node.getNode(r)
		}
		if node == nil {
			node = this.root.getNode(r)
		}

		if node != nil && node.end {
			return true
		}
	}
	return false
}

func (this *TrieFilter) FindFirst(text string) string {
	var node *trieNode
	var tChars = []rune(text)
	var nBuf = this.pool.Get().(*bytes.Buffer)
	defer this.pool.Put(nBuf)

	for _, r := range tChars {
		if this.skip(r) {
			if node == nil {
				nBuf.Reset()
			} else {
				nBuf.WriteRune(r)
			}
			continue
		}

		if node != nil {
			node = node.getNode(r)
		}
		if node == nil {
			nBuf.Reset()
			node = this.root.getNode(r)
		}

		nBuf.WriteRune(r)

		if node != nil && node.end {
			return nBuf.String()
		}
	}

	return ""
}

func (this *TrieFilter) FindAll(text string) []string {
	var node *trieNode
	var tChars = []rune(text)
	var nBuf = this.pool.Get().(*bytes.Buffer)
	defer this.pool.Put(nBuf)
	var nText []string

	for _, r := range tChars {
		if this.skip(r) {
			if node == nil {
				nBuf.Reset()
			} else {
				nBuf.WriteRune(r)
			}
			continue
		}

		if node != nil {
			node = node.getNode(r)
		}
		if node == nil {
			nBuf.Reset()
			node = this.root.getNode(r)
		}

		nBuf.WriteRune(r)

		if node != nil && node.end {
			nText = append(nText, nBuf.String())
			node = nil
			nBuf.Reset()
		}
	}

	return nText
}

func (this *TrieFilter) Replace(text string, replace rune) string {
	var node *trieNode
	var tChars = []rune(text)

	var start = -1
	for i, r := range tChars {
		if this.skip(r) {
			continue
		}

		if node != nil {
			node = node.getNode(r)
		}
		if node == nil {
			start = i
			node = this.root.getNode(r)
		}

		if node != nil && node.end {
			for b := start; b < i+1; b++ {
				tChars[b] = replace
			}
			node = nil
			start = -1
		}
	}

	return string(tChars)
}
