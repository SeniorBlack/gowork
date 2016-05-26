package main

import (
	"fmt"
	"time"
)

func loop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	var ch chan int = make(chan int) //声明channel,并初始化，信道必须用make初始化（分配内存）

	go loop()
	loop()
	time.Sleep(time.Second)
}
