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
	// 创建tar包文件
	fw, err := os.Create("test.tar")
	handleError(err)
	defer fw.Close()

	// 创建tar.Writer
	tw := tar.NewWriter(fw)
	defer tw.Close()

	// 获取要打包的文件的内容
	fr, err := os.Open("test.txt")
	handleError(err)
	defer fr.Close()

	// 获取文件信息
	fi, err := fr.Stat()
	handleError(err)

	// 创建tar.Header
	hdr := new(tar.Header)
	hdr.Name = fi.Name()
	hdr.Size = fi.Size()
	hdr.Mode = int64(fi.Mode())
	hdr.ModTime = fi.ModTime()

	// 写入数据头
	err = tw.WriteHeader(hdr)
	handleError(err)

	// 写入文件数据
	_, err = io.Copy(tw, fr)
	handleError(err)
}
