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
	var runes = []rune(word)

	for _, r := range runes {
		if unicode.IsSpace(r) {
			continue
		}
		r = clearRune(r)

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

func (filter *TrieFilter) Excludes(runes ...rune) {
	for _, r := range runes {
		filter.excludes[clearRune(r)] = struct{}{}
	}
}

func (filter *TrieFilter) Contains(text string) bool {
	var node *trieNode
	var runes = []rune(text)

	for _, r := range runes {
		r = clearRune(r)

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
	var runes = []rune(text)
	var buf = filter.pool.Get().(*bytes.Buffer)
	defer filter.pool.Put(buf)

	for _, r := range runes {
		var nr = clearRune(r)

		if filter.skip(nr) {
			if node != nil {
				buf.WriteRune(r)
			} else {
				buf.Reset()
			}
			continue
		}

		if node != nil {
			node = node.getNode(nr)
		}
		if node == nil {
			buf.Reset()
			node = filter.root.getNode(nr)
		}

		buf.WriteRune(r)

		if node != nil && node.end {
			return buf.String()
		}
	}

	return ""
}

func (filter *TrieFilter) FindAll(text string) []string {
	var node *trieNode
	var runes = []rune(text)
	var buf = filter.pool.Get().(*bytes.Buffer)
	defer filter.pool.Put(buf)
	var texts []string

	for _, r := range runes {
		var nr = clearRune(r)

		if filter.skip(nr) {
			if node != nil {
				buf.WriteRune(r)
			} else {
				buf.Reset()
			}
			continue
		}

		if node != nil {
			node = node.getNode(nr)
		}
		if node == nil {
			buf.Reset()
			node = filter.root.getNode(nr)
		}

		buf.WriteRune(r)

		if node != nil && node.end {
			texts = append(texts, buf.String())
			node = nil
			buf.Reset()
		}
	}

	return texts
}

func (filter *TrieFilter) Replace(text string, replace rune) string {
	var node *trieNode
	var runes = []rune(text)

	var start = -1
	for i, r := range runes {
		r = clearRune(r)

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
				runes[b] = replace
			}
			node = nil
			start = -1
		}
	}

	return string(runes)
}
