package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明了一个带有3个缓存空间的channel
	var channel = make(chan string, 3)

	go func() {
		channel <- "A"
		channel <- "B"
		channel <- "C"
		fmt.Println("我发送了3个数据")
		// 当发送第四个值的时候，goroutine阻塞，所以下面没有执行
		//channel <- "D"
		fmt.Println("我发送了第4个数据")
	}()
	time.Sleep(time.Second)
	receive1:=<-channel
	fmt.Println(receive1)
	receive2:=<-channel
	fmt.Println(receive2)
}