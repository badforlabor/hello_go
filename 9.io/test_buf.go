package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

/*

 */

func testBuf() {
	bb := []byte{}
	//reader := bytes.NewReader(bb)
	buffer := new(bytes.Buffer) // bytes.NewBuffer(bb)
	//b.WriteString("123")
	var a uint16
	a = 1
	b := uint32(2)
	c := uint64(3)
	d := float32(4.1)
	e := float64(5.2)
	s := "12345我"
	fmt.Println([]byte(s))


	//buffer.WriteString(s)
	binary.Write(buffer, binary.LittleEndian, a)
	//binary.Write(buffer, binary.LittleEndian, []byte(s))	// binary不能写slice
	binary.Write(buffer, binary.LittleEndian, b)
	binary.Write(buffer, binary.LittleEndian, c)
	binary.Write(buffer, binary.LittleEndian, d)
	binary.Write(buffer, binary.LittleEndian, e)
	writeString(buffer, s)
	fmt.Println("buf:", buffer.Bytes())

	a = 0
	b = 0
	c = 0
	d = 0
	e = 0
	s = ""

	binary.Read(buffer, binary.LittleEndian, &a)
	//binary.Read(buffer, binary.LittleEndian, tmp)
	binary.Read(buffer, binary.LittleEndian, &b)
	binary.Read(buffer, binary.LittleEndian, &c)
	binary.Read(buffer, binary.LittleEndian, &d)
	binary.Read(buffer, binary.LittleEndian, &e)
	s = readString(buffer)
	fmt.Println(a, b, c, d, e)
	fmt.Println(s, len(bb), buffer.Bytes())

}
func writeString(w io.Writer, str string) {
	n := int32(len(str))
	fmt.Println("write string cnt:", n)
	binary.Write(w, binary.LittleEndian, n)
	w.Write([]byte(str))
}
func readString(r io.Reader) string {
	n := int32(0)
	binary.Read(r, binary.LittleEndian, &n)
	fmt.Println("read string cnt:", n)
	buffer := make([]byte, n)
	r.Read(buffer)
	return string(buffer)
}