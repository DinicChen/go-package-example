package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("12345678"))
	r := bufio.NewReader(rb)
	b, err := r.ReadByte()
	fmt.Printf("%c, %v\n", b, err) // 1, <nil>
}
