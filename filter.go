package sx

type Filter interface {
	// Excludes 用于设置需要排除的内容，主要用于忽略敏感词中间的特殊字符
	// 比如:
	// 将中线(-)设置为排除内容，同时将 "今天" 设置为敏感词
	// 匹配的时候, "今天" 和 "今-天" 都将被判定为敏感词
	Excludes(items ...rune)

	// Contains 用于检测是否有敏感词：如果有敏感词，返回 true；如果没有敏感词，返回 false；
	Contains(text string) bool

	// FindFirst 用于查找出第一个敏感词
	FindFirst(text string) string

	// FindAll 用于查找出所有敏感词
	FindAll(text string) []string

	// ReplaceRune 用于替换敏感词
	Replace(text string, replace rune) string
}
