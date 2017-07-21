package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"bufio"
)

/*
	基本的IO接口
*/

func testInterface() {
	reader := strings.NewReader("测试Reader")
	testReader(reader)

	buf := "测试输出!\n"
	testWriter(os.Stdout, []byte(buf))

	testReaderAt()
	testWriterAt()

	testReadFromWriteTo()

	testSeeker()

	testScanner()
}

func testReader(reader io.Reader) {
	buf := make([]byte, 256)
	n, err := reader.Read(buf)
	if n > 0 {
		fmt.Println(string(buf[0:n]))
	} else if err != nil {
		fmt.Print(err.Error())
	}
}
func testWriter(writer io.Writer, buf []byte) {
	writer.Write(buf)
}
func testReaderAt() {
	reader := strings.NewReader("测试ReaderAt")
	buf := make([]byte, 256)
	n, err := reader.ReadAt(buf, 0)
	if n > 0 {
		fmt.Println(string(buf[0:n]))
	}
	if err != nil {
		fmt.Println(err)
	}
}
func testWriterAt() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteAt([]byte("测试WriteAt"), 0)
}
func testReadFromWriteTo() {
	r := strings.NewReader("测试ReadFromWriteTo")
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(r)
	writer.Flush()

	r.Seek(0, io.SeekStart)
	r.WriteTo(os.Stdout)
}

func testSeeker() {
	r := strings.NewReader("测试ReadFromWriteTo")
	r.Seek(0, io.SeekStart)
	c, _, _ := r.ReadRune()	// 读取一个unicode字符
	b, _ := r.ReadByte()
	fmt.Printf("\nc:%c, byte:%c\n", c, b)
}

func testScanner() {
	/*
		type ByteScanner interface {
			ByteReader
			UnreadByte() error
		}
	*/

	scanner := strings.NewReader("测试scanner")

	c, _, _ := scanner.ReadRune()
	fmt.Printf("scan:%c\n", c)
	c, _, _ = scanner.ReadRune()
	fmt.Printf("scan:%c\n", c)
	scanner.UnreadRune()
	c, _, _ = scanner.ReadRune()
	fmt.Printf("scan:%c\n", c)

	b, _ := scanner.ReadByte()
	scanner.UnreadByte()
	b, _ = scanner.ReadByte()
	fmt.Printf("scan:%c\n", b)
}