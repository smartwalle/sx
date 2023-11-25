package sx

import (
	"bytes"
	"strings"
	"sync"
	"unicode"
)

type trieNode struct {
	end      bool               // 用于标记是否为敏感词结束
	children map[rune]*trieNode // 子节点
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
	}
}

func (node *trieNode) getNode(r rune) *trieNode {
	return node.children[r]
}

type TrieFilter struct {
	pool     *sync.Pool
	root     *trieNode
	excludes map[rune]struct{}
}

func NewTrieFilter(stock WordStock) *TrieFilter {
	var t = &TrieFilter{}
	t.prepare(stock)
	return t
}

func (filter *TrieFilter) prepare(stock WordStock) {
	filter.pool = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
	filter.root = newTrieNode()

	var words = stock.ReadAll()

	for _, word := range words {
		word = strings.TrimSpace(word)
		if len(word) == 0 {
			continue
		}
		filter.addNode(word)
	}
	filter.excludes = make(map[rune]struct{})
	return
}

func (filter *TrieFilter) addNode(word string) {
	var node = filter.root
	var wChars = []rune(word)

	for _, r := range wChars {
		if unicode.IsSpace(r) {
			continue
		}
		r = clear(r)

		if _, ok := node.children[r]; !ok {
			node.children[r] = newTrieNode()
		}
		node = node.children[r]
	}
	node.end = true
}

func (filter *TrieFilter) skip(r rune) bool {
	// 太影响效率
	if /* unicode.IsSpace(r) || unicode.IsPunct(r) || */ filter.inExclude(r) {
		return true
	}
	return false
}

func (filter *TrieFilter) inExclude(r rune) bool {
	_, ok := filter.excludes[r]
	return ok
}

func (filter *TrieFilter) Excludes(items ...rune) {
	for _, item := range items {
		filter.excludes[clear(item)] = struct{}{}
	}
}

func (filter *TrieFilter) Contains(text string) bool {
	var node *trieNode
	var tChars = []rune(text)

	for _, r := range tChars {
		r = clear(r)

		if filter.skip(r) {
			continue
		}

		if node != nil {
			node = node.getNode(r)
		}
		if node == nil {
			node = filter.root.getNode(r)
		}

		if node != nil && node.end {
			return true
		}
	}
	return false
}

func (filter *TrieFilter) FindFirst(text string) string {
	var node *trieNode
	var tChars = []rune(text)
	var nBuf = filter.pool.Get().(*bytes.Buffer)
	defer filter.pool.Put(nBuf)

	for _, r := range tChars {
		var nr = clear(r)

		if filter.skip(nr) {
			if node != nil {
				nBuf.WriteRune(r)
			} else {
				nBuf.Reset()
			}
			continue
		}

		if node != nil {
			node = node.getNode(nr)
		}
		if node == nil {
			nBuf.Reset()
			node = filter.root.getNode(nr)
		}

		nBuf.WriteRune(r)

		if node != nil && node.end {
			return nBuf.String()
		}
	}

	return ""
}

func (filter *TrieFilter) FindAll(text string) []string {
	var node *trieNode
	var tChars = []rune(text)
	var nBuf = filter.pool.Get().(*bytes.Buffer)
	defer filter.pool.Put(nBuf)
	var nText []string

	for _, r := range tChars {
		var nr = clear(r)

		if filter.skip(nr) {
			if node != nil {
				nBuf.WriteRune(r)
			} else {
				nBuf.Reset()
			}
			continue
		}

		if node != nil {
			node = node.getNode(nr)
		}
		if node == nil {
			nBuf.Reset()
			node = filter.root.getNode(nr)
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

func (filter *TrieFilter) Replace(text string, replace rune) string {
	var node *trieNode
	var tChars = []rune(text)

	var start = -1
	for i, r := range tChars {
		r = clear(r)

		if filter.skip(r) {
			continue
		}

		if node != nil {
			node = node.getNode(r)
		}
		if node == nil {
			start = i
			node = filter.root.getNode(r)
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
