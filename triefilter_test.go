package sx_test

import (
	sx "github.com/smartwalle/sx"
	"testing"
)

func getFilter() sx.Filter {
	var stock, _ = sx.NewMemoryStock("福音会", "中国教徒", "统一教", "观音法门", "清海无上师", "盘古", "李洪志", "志洪李", "李宏志", "轮功", "法轮", "轮法功", "三去车仑", "wtf")
	var filter = sx.NewTrieFilter(stock)
	filter.Excludes('-')
	return filter
}

func TestTrieFilter_Contains(t *testing.T) {
	var filter = getFilter()

	var tests = []struct {
		text   string
		expect bool
	}{
		{"福音会", true},
		{"-福音会", true},
		{"--福音会", true},
		{"---福音会", true},
		{"福音会-", true},
		{"福音会--", true},
		{"福音会---", true},
		{"福-音会", true},
		{"福--音会", true},
		{"福-音-会", true},
		{"福--音---会", true},
		{"a福音会", true},
		{"ab福音会", true},
		{"福音会a", true},
		{"福音会ab", true},
		{"三去车仑", true},
		{"wtf 三去车仑", true},
		{"三去车仑 wtf", true},
		{"Golang 的格式化输出和 C 语言的标准输出基本一样。三去车仑", true},
		{"Golang 的格式化输出和 C 语言的标准清海无上师输出基本一样。", true},
		{"T is a type passed to Test functions to manage 观音法门 test state and support formatted test logs.", true},
		{"wtf", true},
		{"WTF", true},
		{"wtf is a type passed to Test", true},
		{"WTf is a type passed to Test", true},
		{"WTF is a type passed to Test", true},
		{"W-TF is a type passed to Test", true},
		{"清海无上师是福音会的中国教徒", true},
		{".福音会", true},
		{"福音会.", true},
		{"福音.会", false},
		{"福-.音会", false},
		{"福.-音会", false},
		{"这是一段干净的文本", false},
	}

	for _, test := range tests {
		if actual := filter.Contains(test.text); actual != test.expect {
			t.Errorf("filter.Contains(%s), 期望得到: %t, 实际得到: %t", test.text, test.expect, actual)
		}
	}
}

func TestTrieFilter_FindFirst(t *testing.T) {
	var filter = getFilter()

	var tests = []struct {
		text   string
		expect string
	}{
		{"福音会", "福音会"},
		{"-福音会", "福音会"},
		{"--福音会", "福音会"},
		{"---福音会", "福音会"},
		{"福音会-", "福音会"},
		{"福音会--", "福音会"},
		{"福音会---", "福音会"},
		{"福-音会", "福-音会"},
		{"福--音会", "福--音会"},
		{"福-音-会", "福-音-会"},
		{"福--音---会", "福--音---会"},
		{"a福音会", "福音会"},
		{"ab福音会", "福音会"},
		{"福音会a", "福音会"},
		{"福音会ab", "福音会"},
		{"三去车仑", "三去车仑"},
		{"wtf 三去车仑", "wtf"},
		{"三去车仑 wtf", "三去车仑"},
		{"Golang 的格式化输出和 C 语言的标准输出基本一样。三去车仑", "三去车仑"},
		{"Golang 的格式化输出和 C 语言的标准清海无上师输出基本一样。", "清海无上师"},
		{"T is a type passed to Test functions to manage 观音法门 test state and support formatted test logs.", "观音法门"},
		{"wtf", "wtf"},
		{"WTF", "WTF"},
		{"wtf is a type passed to Test", "wtf"},
		{"WTf is a type passed to Test", "WTf"},
		{"WTF is a type passed to Test", "WTF"},
		{"W-TF is a type passed to Test", "W-TF"},
		{"清海无上师是福音会的中国教徒", "清海无上师"},
		{".福音会", "福音会"},
		{"福音会.", "福音会"},
		{"福音.会", ""},
		{"福-.音会", ""},
		{"福.-音会", ""},
		{"这是一段干净的文本", ""},
	}

	for _, test := range tests {
		if actual := filter.FindFirst(test.text); actual != test.expect {
			t.Errorf("filter.FindFirst(%s), 期望得到: %s, 实际得到: %s", test.text, test.expect, actual)
		}
	}
}

func TestTrieFilter_FindAll(t *testing.T) {
	var filter = getFilter()

	var tests = []struct {
		text   string
		expect []string
	}{
		{"福音会", []string{"福音会"}},
		{"-福音会", []string{"福音会"}},
		{"--福音会", []string{"福音会"}},
		{"---福音会", []string{"福音会"}},
		{"福音会-", []string{"福音会"}},
		{"福音会--", []string{"福音会"}},
		{"福音会---", []string{"福音会"}},
		{"福-音会", []string{"福-音会"}},
		{"福--音会", []string{"福--音会"}},
		{"福-音-会", []string{"福-音-会"}},
		{"福--音---会", []string{"福--音---会"}},
		{"a福音会", []string{"福音会"}},
		{"ab福音会", []string{"福音会"}},
		{"福音会a", []string{"福音会"}},
		{"福音会ab", []string{"福音会"}},
		{"三去车仑", []string{"三去车仑"}},
		{"wtf 三去车仑", []string{"wtf", "三去车仑"}},
		{"三去车仑 wtf", []string{"三去车仑", "wtf"}},
		{"Golang 的格式化输出和 C 语言的标准输出基本一样。三去车仑", []string{"三去车仑"}},
		{"Golang 的格式化输出和 C 语言的标准清海无上师输出基本一样。", []string{"清海无上师"}},
		{"T is a type passed to Test functions to manage 观音法门 test state and support formatted test logs.", []string{"观音法门"}},
		{"wtf", []string{"wtf"}},
		{"WTF", []string{"WTF"}},
		{"wtf is a type passed to Test", []string{"wtf"}},
		{"WTf is a type passed to Test", []string{"WTf"}},
		{"WTF is a type passed to Test", []string{"WTF"}},
		{"W-TF is a type passed to Test", []string{"W-TF"}},
		{"清海无上师是福音会的中国教徒", []string{"清海无上师", "福音会", "中国教徒"}},
		{".福音会", []string{"福音会"}},
		{"福音会.", []string{"福音会"}},
		{"福音.会", []string{}},
		{"福-.音会", []string{}},
		{"福.-音会", []string{}},
		{"这是一段干净的文本", []string{}},
	}

	for _, test := range tests {
		var actual = filter.FindAll(test.text)
		if len(actual) != len(test.expect) {
			t.Errorf("filter.FindAll(%s), 期望得到: %s, 实际得到: %s", test.text, test.expect, actual)
		}

		for i := 0; i < len(actual); i++ {
			if test.expect[i] != actual[i] {
				t.Errorf("filter.FindAll(%s), 期望得到: %s, 实际得到: %s", test.text, test.expect, actual)
			}
		}
	}
}

func TestTrieFilter_Replace(t *testing.T) {
	var filter = getFilter()

	var tests = []struct {
		text   string
		expect string
	}{
		{"福音会", "***"},
		{"-福音会", "-***"},
		{"--福音会", "--***"},
		{"---福音会", "---***"},
		{"福音会-", "***-"},
		{"福音会--", "***--"},
		{"福音会---", "***---"},
		{"福-音会", "****"},
		{"福--音会", "*****"},
		{"福-音-会", "*****"},
		{"福--音---会", "********"},
		{"a福音会", "a***"},
		{"ab福音会", "ab***"},
		{"福音会a", "***a"},
		{"福音会ab", "***ab"},
		{"三去车仑", "****"},
		{"wtf 三去车仑", "*** ****"},
		{"三去车仑 wtf", "**** ***"},
		{"Golang 的格式化输出和 C 语言的标准输出基本一样。三去车仑", "Golang 的格式化输出和 C 语言的标准输出基本一样。****"},
		{"Golang 的格式化输出和 C 语言的标准清海无上师输出基本一样。", "Golang 的格式化输出和 C 语言的标准*****输出基本一样。"},
		{"T is a type passed to Test functions to manage 观音法门 test state and support formatted test logs.", "T is a type passed to Test functions to manage **** test state and support formatted test logs."},
		{"wtf", "***"},
		{"WTF", "***"},
		{"wtf is a type passed to Test", "*** is a type passed to Test"},
		{"WTf is a type passed to Test", "*** is a type passed to Test"},
		{"WTF is a type passed to Test", "*** is a type passed to Test"},
		{"W-TF is a type passed to Test", "**** is a type passed to Test"},
		{"清海无上师是福音会的中国教徒", "*****是***的****"},
		{".福音会", ".***"},
		{"福音会.", "***."},
		{"福音.会", "福音.会"},
		{"福-.音会", "福-.音会"},
		{"福.-音会", "福.-音会"},
		{"这是一段干净的文本", "这是一段干净的文本"},
	}

	for _, test := range tests {
		if actual := filter.Replace(test.text, '*'); actual != test.expect {
			t.Errorf("filter.Replace(%s), 期望得到: %s, 实际得到: %s", test.text, test.expect, actual)
		}
	}
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
		filter.Replace("台“[中国教徒]”造谣大陆没报道郑州水灾，被连线观众[清海无上师]当场打脸 wtf", '*')
	}
}
