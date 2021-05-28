package main

import "fmt"

var mapvalue map[int]string

func testMap() {

	// map未初始化的时候，为nil。
	// 使用前必须先make初始化
	mapvalue = make(map[int]string)
	mapvalue[1] = "1"

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

type structMap2 struct {
	name string
	age  int
}
func (self *structMap2) modAge(age int) {
	self.age = age
}
func (self structMap2) modAge2(age int) {
	self.age = age
}

func testMap2() {
	m := map[int]structMap2 {
		1: {"user1", 18},
		2: {"user2", 20},
	}

	// 这个操作，并不会修改m
	for _, v := range m {
		v.age += 1
	}
	fmt.Println(m)

	// 重新赋值之后才可以
	for k, v := range m {
		v.age += 1
		m[k] = v
	}
	fmt.Println(m)

	// m[3].modAge(3)	// compile err
	m[3].modAge2(3)
	fmt.Println(m[3].age)

	// 可以直接取
	var old = m[4]
	old.modAge(4)
	m[4] = old
	fmt.Println(m[4])

	// 此时 ok=false
	var old2, ok = m[5]
	fmt.Println(ok, old2.age)
}
