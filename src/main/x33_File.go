package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("d:/test.txt")

	if err != nil {
		fmt.Println("打开错误")
		return
	}

	defer file.Close()

	//读取文件内容并输出
	reader := bufio.NewReader(file)
	for {
		res, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(res)
	}

}
