package main

import (
	"fmt"
	"sync"
)

// 声明一个map，用于存储每个数的阶乘
//var myRes = make(map[int]int,100)//map是映射，未学（建立事物关联的容器）
var (myRes = make(map[int]int)
	// 声明全局互斥锁
	lock sync.Mutex)

func main() {

	// 启用100个goroutine同时求1-100以内各个数的阶乘
	for i := 1; i <=100; i++ {
		go factorial(i)

	}
	fmt.Println(myRes)

}

// 求n的阶乘，并将结果写进myRes
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *=i
	}

	lock.Lock()   // 在进行存储操作之前，先加锁
	myRes[n] = res
	lock.Unlock()  // 当存储完毕后，进行解锁

}