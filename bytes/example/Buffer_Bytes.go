package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("0123456"))
	unread := b.Bytes()
	fmt.Println(string(unread)) // 0123456
	for i, c := range unread {
		unread[i] = 'A' + c - '0'
	}

	var buff [7]byte
	b.Read(buff[:])
	fmt.Println(string(buff[:])) // ABCDEFG
}
