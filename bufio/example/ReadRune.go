package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	b := []byte("你好世界")
	rb := bytes.NewBuffer(b)
	r := bufio.NewReader(rb)
	c, size, err := r.ReadRune()
	if err == nil {
		fmt.Println(string(c))          // 你
		fmt.Printf("%x, %d\n", c, size) // 4f60, 3
		fmt.Printf("%x\n", b[:size])    // e4bda0
	}
}
