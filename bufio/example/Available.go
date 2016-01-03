package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	fmt.Println(w.Available()) // 4096
	w.WriteByte('1')
	fmt.Println(w.Available()) // 4095
}
