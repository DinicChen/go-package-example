package main

import (
	"archive/tar"
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
	// 打开tar包文件
	fr, err := os.Open("test.tar")
	handleError(err)
	defer fr.Close()

	// 创建tar.Reader
	tr := tar.NewReader(fr)

	for {
		// 获取下一个文件，第一个文件也用此方法获取
		hdr, err := tr.Next()
		// 已读到文件尾
		if err == io.EOF {
			break
		}
		handleError(err)

		// 通过创建文件获得*io.Writer
		fw, _ := os.Create(hdr.Name)
		handleError(err)

		// 拷贝数据
		_, err = io.Copy(fw, tr)
		handleError(err)
	}
}
