package main

import (
	"fmt"
	"time"
)

/*
	channel用法
	defer
	coroutine用法(go关键字）
	select用法

	详解：
		channel等价于生产者消费者模型
			chan_var <- 1		// 相当于，生产了一个数字1
			a := <- chan_var	// 相当于，消费了管道，如果管道里面没有内容，那么会一直阻塞。
		select 用来解决channel阻塞问题。
			从多个channel中选择一个有数据的执行，如果没有可用的，执行default
			如果没有defalut，那么会阻塞住。

		funcb() {defer funca();...} funcb函数体中的内容都执行完毕后，再执行funca

		go funca();		另开一个线程执行funca。

*/

func main() {
	//testBaseChannel()
	testBaseChannel2()
	//testChannelRange()
	//testChannelRangeSelect()
	//testSelect()
}
func testBaseChannel() {
	data := make(chan int)

	go func() {
		// 如果把这里注释掉，会因为没有任何消费data的地方，会直接阻塞函数所在线程
		a := <-data
		fmt.Println("消费：", a)
	}()

	data <- 5
	fmt.Println("end test base channel")
}
func testBaseChannel2() {
	data := make(chan int)
	data1 := make(chan int)

	go func() {
		// 如果把这里注释掉，会因为没有任何消费data的地方，会直接阻塞函数所在线程
		a := <-data
		fmt.Println("消费：", a)

		time.Sleep(3 * time.Second)

		a = <-data1
		fmt.Println("消费：", a)
	}()

	fmt.Println("生成数据5")
	data <- 5
	fmt.Println("已经消费了data数据")

	fmt.Println("生成数据4")
	// 由于放入channle中的4没有立即消费掉，会导致3秒后才出现“已经消费了data1数据”字符串
	data1 <- 4
	fmt.Println("已经消费了data1数据")

	fmt.Println("end test base channel")
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
		exit <- false
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
	a := <-exit
	fmt.Println("exit:", a)
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
			case q = <-quit:
				//q = true
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

	// 如果select里面不加任何内容，会直接导致死锁。
	// select{}
}

func testSelect() {
	data := make(chan int)
	go func() {
		var a int
		select {
			case a = <- data:
				fmt.Println("a:", a)
		}
		// 上面的select没有default，会一直阻塞，知道data中有数据
		fmt.Println("select blocked.")
	}()

	time.Sleep(3 * time.Second)
	data <- 5
	fmt.Println("pre end")
	time.Sleep(5 * time.Second)
	fmt.Println("end")
}