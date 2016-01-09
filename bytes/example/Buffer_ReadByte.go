package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("123456"))
	c, err := b.ReadByte()
	fmt.Println(c, err)            // 49 <nil>
	fmt.Println(string(b.Bytes())) // 23456
}
