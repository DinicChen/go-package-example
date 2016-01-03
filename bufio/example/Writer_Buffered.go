package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	fmt.Println(w.Buffered()) // 0
	w.WriteByte('1')
	fmt.Println(w.Buffered()) // 1
	w.Flush()
	fmt.Println(w.Buffered()) // 0
}
