package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("1234,56$"))
	r := bufio.NewReader(rb)
	line, _ := r.ReadSlice(',')
	fmt.Println(string(line)) // 1234,

	// unread ','
	fmt.Println(r.UnreadByte()) // <nil>
	line, _ = r.ReadSlice('$')
	fmt.Println(string(line)) // ,56$
}
