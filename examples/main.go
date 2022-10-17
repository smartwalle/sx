package main

import (
	"fmt"
	"github.com/smartwalle/sx"
)

func main() {
	var fs, _ = sx.NewFileStock("./ck.txt")
	var filter = sx.NewTrieFilter(fs)
	filter.Excludes('-', ' ')
	fmt.Println(filter.Contains("sss福音会"))
	fmt.Println(filter.Contains("福音会"))
	fmt.Println(filter.Contains("s刘广智空军sss福.音会"))
	fmt.Println(filter.FindFirst("福福音会"))
	fmt.Println(filter.FindFirst("福-音会"))
	fmt.Println(filter.FindFirst("福音会ss"))
	fmt.Println(filter.FindFirst("s福音会"))
	fmt.Println(filter.FindAll("sss刘广智空军sss福音会"))
	fmt.Println(filter.FindAll("sstrychnine福音会"))
	fmt.Println(filter.Replace("s刘广智空军sss福音会", '*'))
	fmt.Println(filter.Replace("刘广智空军sss福音会", '*'))
	fmt.Println(filter.Replace("world刘广智-空军hello福音会", '*'))
	fmt.Println(filter.Replace("hello", '*'))
	fmt.Println(filter.Replace("w", '*'))
}
