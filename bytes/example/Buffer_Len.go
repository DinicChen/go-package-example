package main

import (
	    "bytes"
		    "fmt"
		)

		func main() {
			    b := bytes.NewBuffer([]byte("123456"))
				    fmt.Println(b.Len())
					    var buff [3]byte
						    b.Read(buff[:])
							    fmt.Println(b.Len())
							}
							代码输出

							6
							3
