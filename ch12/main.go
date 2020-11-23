package main

import (
	"fmt"
	"strconv"
	"time"
)

func producer(i chan int) {
	for {
		select {
		case a := <-i:
			fmt.Println("接收--->" + strconv.Itoa(a))
		}
	}
}

func consumer(i chan int) {
	for j := 0; j < 10; j++ {
		fmt.Println("发送--->" + strconv.Itoa(j))
		i <- j
	}
}

func main() {
	ch := make(chan int, 2)
	go producer(ch)
	go consumer(ch)
	time.Sleep(10 * time.Second)
}
