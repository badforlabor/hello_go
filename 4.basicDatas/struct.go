package main

import "fmt"

type Resource struct {
	id   int
	name string
}

func testStruct() {
	type Classify struct {
		id int
	}
	type Other struct {
		other string
	}
	type User struct {
		Resource // 这种叫做匿名字段
		Classify
		name string
		Other
	}
	user := User{
		Resource{1, "people"},
		Classify{100},
		"jack",
		Other{"other"},
	}
	fmt.Printf("%s %d %s %d %s\n",
		user.name, user.Resource.id, user.Resource.name,
		// 名字有重叠的，就得显示调用了
		user.Classify.id,
		// 下面这个由于Other和User没有字段重叠，所以可以直接用user.other。
		user.other)

	type User2 struct {
		*Resource
		// 与上面的指针重复了，只能显示生命
		// Resource
		r    Resource
		name string
	}

	r := Resource{}
	proc(r)
	fmt.Println(r)
	proc2(&r)
	fmt.Println(r)
}

// 传值，复制一份
func proc(r Resource) {
	r.id = 1
}

// 传引用才好用。
func proc2(r *Resource) {
	r.id = 1
}
