package main

import (
	"fmt"
	"time"
)

/*
	channel用法
*/

func main() {
	//testChannelRange()
	testChannelRangeSelect()
}

func testChannelRange() {

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

func testChannelRangeSelect() {
	defer func() {
		fmt.Println("exit.")
	}()

	data := make(chan int)
	exit := make(chan bool)
	quit := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		q := false
		lastTime := time.Now()
		for !q {

			empty := false
			for !empty {
				select {
				case d := <-data:
					fmt.Println("data:", d)
				default:
					empty = true
				}
			}

			select {
			case <-quit:
				q = true
			default:
			}

			tNow := time.Now()
			frame := time.Second / 30
			if tNow.Sub(lastTime) < frame {
				fmt.Println("time:", time.Now().String())
				time.Sleep(frame - tNow.Sub(lastTime))
			}
			fmt.Println("time delta:", tNow.Sub(lastTime))
			lastTime = tNow
		}
		fmt.Println("recv over.")
		exit <- true
	}()
	go func() {
		for false {
			fmt.Println("echo.")
			time.Sleep(1 * time.Second)
		}
	}()

	data <- 1
	data <- 2
	data <- 3
	quit <- true // 执行这条语句之后，会让select成功，然后整个进程退出
	//close(data)

	fmt.Println("send over.")
	<-exit
}
