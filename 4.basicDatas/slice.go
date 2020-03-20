package main

import (
	"fmt"
)

/*
	slice实际上是一个结构体，指向数组的结构体！
	golang内部实现是：
		struct Slice
		{
			byte* array;
			uintgo len;
			uintgo cap;
		}
	append函数可以追加数据，如果没有超过原数据大小，那么相当于修改，如果超过了，那么重新申请内存，在新内存上修改
	copy(dst, src)函数拷贝数据
*/

func testSlice() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := data[1:4:5] // 取[1,4-1]数据。[low:high:max], len=high-low, cap=max-low
	s2 := data[1:4]   // low:hi
	fmt.Println("s1:", s1)
	fmt.Println(s2)
	s2[0] = 100 // 相当于data[2] = 100
	fmt.Println(data)

	// 各种简写模式
	s3 := data[:4]   // 相当于0:4
	s4 := data[:4:5] // 相当于0:4:5
	s5 := data[1:]
	s6 := data[:]
	fmt.Println(s3, s4, s5, s6)

	// 通过append追加数据
	s7 := data[:]
	s8 := append(s7, 100, 200) // data不变，s8新申请了一块内存，复制data数据，然后再追加两个数据
	fmt.Println(s8, data)
	fmt.Printf("%p %p \n", &s8[0], data)

	s9 := data[:3]
	s10 := append(s9, 100, 200) // 相当于data[3]=100,data[4]=100
	fmt.Println(s9, s10, data)
	fmt.Printf("%p %p %p\n", &s9[0], data, &s10[0])

	// copy数据
	{
		data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		s := data[8:]
		s2 := data[:5]
		copy(s2, s)
		fmt.Println(s, s2)
	}

	fmt.Printf("addr:%p \n", data)
	testSliceAddr(data)
	testSliceAddr(s1)

	// 切片和数据传参的区别
	{
		// 这个是数组，执行changeValue2后，并不会修改原有数组的值
		data := [5]int{0, 1, 2, 3, 4}
		changeValue2(data)
		fmt.Println("changevalue2:", data)
		s1 := data[0:2]
		//changeValue(data)
		//fmt.Println("changevalue:", data)

		// 这个是切片，内部引用的是数组，执行changeValue后，数值会改变
		changeValue(s1)
		fmt.Println("changevalue:", data)
		d1 := []int{0, 1, 2, 3, 4}
		changeValue(d1)
		fmt.Println(d1)

		appendValue(d1)
		fmt.Println(d1)

		appendValuePtr(&d1)
		fmt.Println(d1)
	}
}
func changeValue2(d [5]int) {
	d[0] = 100
}
func changeValue(d []int) {
	d[0] = 100
}
func testSliceAddr(s []int) {
	fmt.Printf("%p \n", s)
}
func appendValue(d []int) {
	// 无效果
	// 新开辟了一块内存
	d = append(d, 0, 0, 0)
}
func appendValuePtr(d *[]int) {
	// 有效果
	// 新开辟了内存，但是又被原有的d指向了。
	*d = append(*d, 0, 0, 0)
}


func testSlice2() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 这个操作并不会导致结果变成1,2,3,4...
	for _, v := range data {
		v += 1
	}
	fmt.Println(data)

	for i, _ := range data {
		data[i] += 1
	}
	fmt.Println(data)
}