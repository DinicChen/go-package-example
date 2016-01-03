package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("123456"))
	r := bufio.NewReader(rb)

	r.ReadByte()
	fmt.Println(r.UnreadRune()) // bufio: invalid use of UnreadRune

	c, _, _ := r.ReadRune()
	fmt.Printf("read %s\n", string(c)) // read 2
	fmt.Println(r.UnreadRune())        // <nil>

	c, _, _ = r.ReadRune()
	fmt.Printf("read %s\n", string(c)) // read 2
}
