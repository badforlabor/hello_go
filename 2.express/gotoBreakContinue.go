package main

import "fmt"

/*
	switch，不需要写break，自动break
	如果想不break，那么用fallthrough关键字
	可以忽略条件表达式
*/
func testOther() {

	x := []int{1, 2}
	for _, i := range x {

		// 当i=1的时候，看到的结果是1
		// 当i=2的时候，看到的结果是2 3
		switch i {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
			fallthrough // 继续下一条
		case 3:
			fmt.Println("3")
		}

	}

	for _, i := range x {
		// 忽略条件表达式的switch，就相当于if else
		switch {
		case i == 1:
			fmt.Println("1")
		case i == 2:
			fmt.Println("2")
			fallthrough // 继续下一条
		case i == 3:
			fmt.Println("3")
		}
	}

	for _, i := range x {
		// 忽略条件表达式的switch，就相当于if else
		switch {
		case i == 1:
			fmt.Println("1")
		case i == 2:

			fmt.Println("2")
			fallthrough // 继续下一条
		case i == 3:
			// 当然也可以用break
			if i == 2 {
				break
			}
			fmt.Println("3")
		}
	}
}
