package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)
//type Reader interface {
//	Read(p []byte) (n int, err error)
//}
func main() {
	reader := strings.NewReader("Clear is better than clever")//通过 string.NewReader(string) 创建一个字符串读取器
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil{
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}