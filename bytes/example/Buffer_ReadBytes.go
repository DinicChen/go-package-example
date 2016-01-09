package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("123,456"))
	line, err := b.ReadBytes(',')
	fmt.Println(string(line)) // 123,
	fmt.Println(err)          // <nil>
}
