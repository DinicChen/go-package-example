package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("a string to be read"))
	wb := bytes.NewBuffer(nil)
	r := bufio.NewReader(rb)
	w := bufio.NewWriter(wb)
	rw := bufio.NewReadWriter(r, w)

	// use rw to read
	var rbuf [128]byte
	if n, err := rw.Read(rbuf[:]); err != nil {
		return
	} else {
		fmt.Println(string(rbuf[:n])) // a string to be read
	}

	// use rw to write
	if _, err := rw.Write([]byte("a string to be written")); err != nil {
		return
	}
	rw.Flush()
	fmt.Println(string(wb.Bytes())) // a string to be written
}
