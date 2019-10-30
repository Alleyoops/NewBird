package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input:")
	str, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println(str)
	}
}