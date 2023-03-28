package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
)

func main() {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	}{
		{"Golang.txt", "http://c.biancheng.net/golang/"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}

	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		fmt.Println(err)
	}
	// 将压缩文档内容写入文件
	f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f)

}
