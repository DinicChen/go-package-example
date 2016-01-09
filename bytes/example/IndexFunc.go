package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("你好世界")

	fmt.Println(bytes.IndexFunc(s, func(r rune) bool {
		return r == '好'
	})) // 3

	fmt.Println(bytes.IndexFunc(s, func(r rune) bool {
		return r == '!'
	})) // -1
}
