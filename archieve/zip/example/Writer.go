package main

import (
	"archive/zip"
	"fmt"
	"os"
)

// 处理错误
func handleError(err error) {
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
}

func main() {
	// 创建zip包文件
	fw, err := os.Create("test.zip")
	handleError(err)
	defer fw.Close()

	// 创建zip.Writer
	zw := zip.NewWriter(fw)
	defer zw.Close()

	// 获取要打包的文件的内容
	fr, err := os.Open("test.txt")
	handleError(err)
	defer fr.Close()

	// 获取文件信息
	fi, err := fr.Stat()
	handleError(err)

	// 创建zip.FileHeader
	fh := new(zip.FileHeader)
	fh.Name = fi.Name()
	fh.UncompressedSize = uint32(fi.Size())
	fw2, err := zw.CreateHeader(fh)
	handleError(err)

	// 读取文件数据
	buf := make([]byte, fi.Size())
	_, err = fr.Read(buf)
	handleError(err)

	// 写入数据到zip包
	_, err = fw2.Write(buf)
	handleError(err)
}
