package main

import (
	"fmt"
	"os"
)

func main() {
	//create and close
	fp, err := os.Create("proverb.txt")
	if err != nil {
		fmt.Println("Create failed!")
	}
	fp.Write([]byte("don't communicate by sharing memory share memory by communicating"))
	fp.Close()
	//open and read then close
	fp, err = os.Open("proverb.txt")
	if err != nil {
		fmt.Println("open failed!")
	}
	read := make([]byte, 100)
	fp.Read(read)
	fmt.Println(string(read[:]))

}
