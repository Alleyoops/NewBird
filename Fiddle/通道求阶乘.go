package main

import (
	"fmt"
	"time"
)

var (
	ch    = make(chan int, 20)
	myres = [20]int{}
)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch <- res
	myres[n-1] = <-ch
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	time.Sleep(time.Second)
	for k, v := range myres {
		fmt.Println(k+1, ":", v)
	}

}
