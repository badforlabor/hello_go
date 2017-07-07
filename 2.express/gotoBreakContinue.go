package main

import "fmt"

/*
	goto,
	break   可以用于for，switch，select
	continue   只能用于for
*/
func testOther() {

	// continue
	for i := 0; i < 10; i++ {
		if i < 9 {
			// 跳过前就此
			continue
		}
		// 直接终止for循环
		break
		fmt.Println("10000000000000")
	}

	var it int
	for {
		fmt.Println("gooooooooooo")
		it++
		if it > 0 {
			goto BREAK
		}
	}
BREAK:
	fmt.Println("BREAK")

	// break, continue 还能这样用
L1:
	for i := 0; i < 3; i++ {
	L2:
		for j := 0; j < 3; j++ {
			if j > 2 {
				continue L2
			}
			if i > 1 {
				break L1
			}
			fmt.Println("i:", i, "j:", j)
		}
	}
}
