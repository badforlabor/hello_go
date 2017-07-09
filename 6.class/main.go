package main

import (
	"fmt"
	"unsafe"
)

/*
	golang中没有class关键字
	没有继承，没有多态
	内存布局和C struct相同，没有任何附加的object信息（c++中的多态信息）
*/

type User struct {
	id   int
	name string
}
type Manager struct {
	User
	title string
}

// 类的方法，也可以是这种形式：func (m Manager) memoryInfo() {} 不过，这样会复制一份Manager，不推荐！
// 而且类的方法，跟普通的方法没啥区别，比如func memoryInfo(m *Manager) {}
func (m *Manager) memoryInfo() {
	// 如果直接unsafe.Sizeof(m)，那么会返回sizeof(ptr)，得用unsafe.Sizeof(*m)
	fmt.Printf("m:\t%p, size:%d, align:%d\n", m, unsafe.Sizeof(m), unsafe.Alignof(m))
	fmt.Printf("m:\t%p, size:%d, align:%d\n", m, unsafe.Sizeof(*m), unsafe.Alignof(*m))
	fmt.Printf("m.id\t%p, offset:%d\n", &m.id, unsafe.Offsetof(m.id))
	fmt.Printf("m.name\t%p, offset:%d\n", &m.name, unsafe.Offsetof(m.name))
	fmt.Printf("m.title\t%p, offset:%d\n", &m.title, unsafe.Offsetof(m.title))
}
func (u *User) info() {
	fmt.Printf("%d,%s,%p\n", u.id, u.name, u)
}

func main() {
	fmt.Println("test end")

	m := Manager{}
	m.memoryInfo()

	// 支持匿名字段的函数，哈哈，这个奇葩啊，以下相当于m.User.info()
	m.info()
	m.User.info()
}
