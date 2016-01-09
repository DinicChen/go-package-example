package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("123456"))
	n, err := b.Read(nil)
	fmt.Printf("%d, %v\n", n, err) // 0, <nil>
	var buff [6]byte
	n, err = b.Read(buff[:])
	fmt.Printf("%d, %v, %s\n", n, err, buff[:n]) // 6, <nil>, 123456
	n, err = b.Read(buff[:])
	fmt.Printf("%d, %v, %s\n", n, err, buff[:n]) // 0, EOF,
}
