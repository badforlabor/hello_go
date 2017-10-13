package main

import (
	"sort"
	"fmt"
)

func main() {
	testSortString()
	testSortByFunc()
	testMySearch()
}

func testSortString() {
	// Sort methods are specific to the builtin type;
	// here's an example for strings. Note that sorting is
	// in-place, so it changes the given slice and doesn't
	// return a new one.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 反转
	{
		sort.Sort(sort.Reverse(sort.StringSlice(strs)))
		fmt.Println("Strings:", strs)
	}

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

//==============================================================================================
// 实现sort.Interface即可调用sort.Sort了。
type ByLength []string
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
func testSortByFunc() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

	// 反转
	{
		sort.Sort(sort.Reverse(ByLength(fruits)))
		fmt.Println(fruits)
	}
}

//==============================================================================================
// 测试查找
func testSearch() {
	// 默认的search函数，得先排序数据之后，才能查找
	data := []int{27, 15, 8, 9, 12, 4, 17, 19, 21, 23, 25}
	sort.Ints(data)
	idx := sort.Search(len(data), func(idx int) bool {
		if data[idx] == 10 {
			return true
		}
		return false
	})
	fmt.Println(idx)
}

//==============================================================================================
// 自定义查找
// 查找int，返回索引
func IntSearch(datas []int, v int) int {
	for i, v1 := range datas {
		if v1 == v {
			return i
		}
	}
	return -1
}
func SliceSearch(length int, callback func (int) bool) int {
	for i:=0; i<length; i++ {
		if callback(i) {
			return i
		}
	}
	return -1
}
func testMySearch() {
	{
		data := []int{27, 15, 8, 9, 12, 4, 17, 19, 21, 23, 25}
		idx := IntSearch(data, 23)
		fmt.Println(idx)
	}
	{
		data := []int{27, 15, 8, 9, 12, 4, 17, 19, 21, 23, 25}
		idx := SliceSearch(len(data), func(ii int) bool {
			return data[ii] == 23
		})
		fmt.Println(idx)
	}
}

