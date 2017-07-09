package main

import (
	"fmt"
	"time"
)

/*
	channel用法
*/

func main() {

	defer func() {
		fmt.Println("exit.")
	}()

	data := make(chan int)
	exit := make(chan bool)
	go func() {
		for d := range data {
			fmt.Println(d)
		}
		fmt.Println("recv over.")
		exit <- true
	}()
	go func() {
		for {
			fmt.Println("echo.")
			time.Sleep(1 * time.Second)
		}
	}()

	data <- 1
	data <- 2
	data <- 3
	close(data)

	fmt.Println("send over.")
	<-exit
}
