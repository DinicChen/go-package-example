package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("123456"))
	fmt.Println(string(b.Next(4))) // 1234
	fmt.Println(string(b.Next(4))) // 56
}
