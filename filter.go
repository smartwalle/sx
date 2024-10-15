package sx

type Filter interface {
	// Excludes 用于设置需要排除的内容，主要用于忽略敏感词中间的特殊字符
	// 比如:
	// 将中线(-)设置为排除内容，同时将 "今天" 设置为敏感词
	// 匹配的时候, "今天" 和 "今-天" 都将被判定为敏感词
	Excludes(runes ...rune)

	// Contains 用于检测是否有敏感词：如果有敏感词，返回 true；如果没有敏感词，返回 false；
	Contains(text string) bool

	// FindFirst 用于查找出第一个敏感词
	FindFirst(text string) string

	// FindAll 用于查找出所有敏感词
	FindAll(text string) []string

	// Replace 用于替换敏感词
	Replace(text string, replace rune) string
}

func clearRune(r rune) rune {
	if r >= 65 && r <= 90 {
		// 大写字母转换成小写字母
		r += 32
	} else if r >= 65313 && r <= 65338 {
		// 大写英文全角转换成小写字母
		r = 97 + (r - 65313)
	} else if r >= 65345 && r <= 65370 {
		// 小写英文全角转换成小写字母
		r = 97 + (r - 65345)
	} else if r >= 65280 && r <= 65375 {
		// 其它全角转换成半角
		r = 32 + (r - 65280)
	} else {
		//8216: 39, // '
		//8217: 39, // '
		//8220: 34, // "
		//8221: 34, // "
		//12290: 46, // .
		//12304: 91, // [
		//12305: 93, // ]
		//65281: 33, // !
		//65288: 40, // (
		//65289: 41, // )
		//65292: 44, // ,
		//65306: 58, // :
		//65307: 59, // ;
		//65311: 63, // ?
	}
	return r
}
