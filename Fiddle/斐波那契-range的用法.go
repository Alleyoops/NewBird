package main

import "fmt"

func main() {
	go animation()
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func animation() {
	for {
		for _,r := range  "qedwf" {//下划线表示字符串里元素的索引
			fmt.Printf("\r%", r)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)//函数的递归，效率低
}