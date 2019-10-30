package main

import (
	"fmt"
	"time"
)

func main() {
	var channel = make(chan int)//channel的类型为chan int
	var send = 6666
	// 关键字go后跟的是需要被并发执行的代码块，它由一个匿名函数代表
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以
	go func() {
		channel <- send//send为int 将其值传到通道channel
		fmt.Println("数据已发送")
	}()//至于为什么有空括号，看匿名函数
	// 这里让主线程休眠1秒钟
	// 以使上面的goroutine有机会被执行
	time.Sleep(1 * time.Second)//是主协程休眠一秒

	//没有以下代码，通道的数据不会被取出，造成堵塞，所以上面的print不会执行(因为是不带缓存的通道)
	//receive:=<-channel//将其取出
	//fmt.Println(receive)
}