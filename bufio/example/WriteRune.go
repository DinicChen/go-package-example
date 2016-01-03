package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wb := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wb)
	n, err := w.WriteRune('你')
	if err != nil {
		return
	}

	w.Flush()
	fmt.Println(n)          // 3
	fmt.Println(wb.Bytes()) // [228 189 160]
}
