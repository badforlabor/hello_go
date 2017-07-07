package main

import (
	"fmt"
	"math"
	"runtime"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

var c, python, java bool

func main() {
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(add2(1, 2))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var i int
	fmt.Println(i, c, python, java, Pi)

	main2()
	main3()
	main4()
	main5()
	main7()
	main8()
	main6()
	main9()
	main10()
	main11()
	main12()
	main13()
	main14()
	testVertex()
	testMyFloat()
	testInterface()
	test_routine()
	testChannel()
}

func main2() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func main3() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	// var c, python, java := true, false, "no!" // 错误，var不能和 := 一起使用

	fmt.Println(i, j, k, c, python, java)
}

const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main4() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func main5() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main6() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
func main7() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
func main8() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main9() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])
}
func main10() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
func main11() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main12() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main13() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)

}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main14() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
