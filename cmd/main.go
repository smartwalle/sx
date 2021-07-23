package main

import (
	"fmt"
	sx "github.com/smartwalle/dbm"
)

func main() {
	var fs, _ = sx.NewFileStock("./ck.txt")
	var filter = sx.NewTrieFilter(fs)
	filter.Excludes('-', ' ')
	fmt.Println(filter.Contains("sss福音会"))
	fmt.Println(filter.Contains("福音会"))
	fmt.Println(filter.Contains("s刘广智空军sss福.音会"))
	fmt.Println(filter.FindFirst("福福音会"))  // bug
	fmt.Println(filter.FindFirst("福-音会"))  // bug
	fmt.Println(filter.FindFirst("福音会ss")) // bug
	fmt.Println(filter.FindFirst("s福音会"))  // bug
	fmt.Println(filter.FindAll("sss刘广智空军sss福音会"))
	fmt.Println(filter.FindAll("sstrychnine福音会"))
	fmt.Println(filter.ReplaceRune("s刘广智空军sss福音会", '*'))
	fmt.Println(filter.ReplaceRune("刘广智空军sss福音会", '*'))
	fmt.Println(filter.ReplaceRune("world刘广智-空军hello福音会", '*'))
	fmt.Println(filter.ReplaceRune("hello", '*'))
	fmt.Println(filter.ReplaceRune("w", '*'))
}
