package main

import "fmt"

func main() {
	send := make(chan int)
	go counter(send)
	squarer(send)
}

func counter(out chan<- int) {//chan<- int只发送给通道
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int) {//<-chan int只接收从通道
	for v := range in {
		fmt.Println(v)
	}

}