package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "./output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v", err)
		return
	}
	defer file.Close()
	str := "http://c.biancheng.net/golang/\n" // \n\r表示换行  txt文件要看到换行效果要用 \r\n
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()

	file1, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("文件打开失败 = ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file1)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束...")
}
