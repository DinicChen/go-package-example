package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// 处理错误
func handleError(err error) {
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
}

func main() {
	// 打开zip包文件
	rc, err := zip.OpenReader("test.zip")
	handleError(err)

	// 逐个读取zip包内的单独文件
	for _, f := range rc.File {
		// 打开包中的文件
		r, err := f.Open()
		handleError(err)

		// 将包中的文件数据写出
		fw, _ := os.Create(f.FileInfo().Name())
		handleError(err)

		// 拷贝数据
		_, err = io.Copy(fw, r)
		handleError(err)
	}
}
