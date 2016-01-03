package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	w.WriteByte('a')
	w.Flush()
	fmt.Println(string(wb.Bytes())) // a
}
