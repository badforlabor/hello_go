package main

import "fmt"

func testMap() {
	{
		m := map[int]struct {
			name string
			age  int
		}{
			1: {"user1", 18},
			2: {"user2", 20},
		}
		fmt.Println(m)

		// map取到的值是复制品！修改之后，没有任何效果
		u := m[0]
		u.age = 19
		fmt.Println(m)
	}
	{
		m := make(map[string]int, 1000)

		// 判定map中是否有某个元素
		if v, ok := m["a"]; ok {
			fmt.Println(v)
		}

		// 设置map的值
		m["a"] = 100
		m["a"] = 101

		// 删除值
		delete(m, "c")
	}
}
