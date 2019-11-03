package main

import (
	"fmt"
	"time"
)

func hey(a int) {
	for i := 2; i <= a; i++ {
		var j int
		for j = 2; j <= i; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == i {
			fmt.Println(i)
		}
	}
}

func main() {
	go hey(10000)
	time.Sleep(time.Second)
}
