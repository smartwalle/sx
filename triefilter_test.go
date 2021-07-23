package sx_test

import (
	sx "github.com/smartwalle/dbm"
	"testing"
)

func BenchmarkTrieFilter_Contains(b *testing.B) {
	var fs, _ = sx.NewFileStock("./ck.txt")
	var filter = sx.NewTrieFilter(fs)
	filter.Excludes('-')
	b.Log(filter.Contains("sss福音会"))
}

func BenchmarkTrieFilter_FindFirst(b *testing.B) {
}

func BenchmarkTrieFilter_FindAll(b *testing.B) {
}

func BenchmarkTrieFilter_ReplaceRune(b *testing.B) {
}
