package sx_test

import (
	sx "github.com/smartwalle/sx"
	"testing"
)

func getFilter() sx.Filter {
	var stock, _ = sx.NewMemoryStock("福音会", "中国教徒", "统一教", "观音法门", "清海无上师", "盘古", "李洪志", "志洪李", "李宏志", "轮功", "法轮", "轮法功", "三去车仑")
	var filter = sx.NewTrieFilter(stock)
	filter.Excludes('-')
	return filter
}

func BenchmarkTrieFilter_Contains(b *testing.B) {
	var filter = getFilter()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		filter.Contains("台“[中国教徒]”造谣大陆没报道郑州水灾，被连线观众[清海无上师]当场打脸")
	}
}

func BenchmarkTrieFilter_FindFirst(b *testing.B) {
	var filter = getFilter()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		filter.FindFirst("台“[中国教徒]”造谣大陆没报道郑州水灾，被连线观众[清海无上师]当场打脸")
	}
}

func BenchmarkTrieFilter_FindAll(b *testing.B) {
	var filter = getFilter()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		filter.FindAll("台“[中国教徒]”造谣大陆没报道郑州水灾，被连线观众[清海无上师]当场打脸")
	}
}

func BenchmarkTrieFilter_ReplaceRune(b *testing.B) {
	var filter = getFilter()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		filter.Replace("台“[中国教徒]”造谣大陆没报道郑州水灾，被连线观众[清海无上师]当场打脸", '*')
	}
}
