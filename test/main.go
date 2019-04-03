package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a = "Hello 中文  你怎么劝你"
	fmt.Println(len(a))
	fmt.Println(utf8.RuneCountInString(a))
	b := []rune(a)
	fmt.Println(string(b[0:12]))
	fmt.Println(a[0:28])
}
