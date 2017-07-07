package main

/*
	if语句

*/
func testIf() {
	b := false

	// 不需要这么写 if(b)
	// 大括号不能省略
	if b {
		b = false
	}

	// 也可以这么写
	if c := b; c {
		c = false
	}

	// 这样写是错误的，不能出现 逗号
	/*
		if d := b, e := d; d {
			d = b
			e = b
		}
	*/
}
