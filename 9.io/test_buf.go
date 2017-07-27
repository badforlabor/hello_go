package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*

 */

func testBuf() {
	b := []byte{}
	//reader := bytes.NewReader(b)
	writer := new(bytes.Buffer) // bytes.NewBuffer(b)
	//b.WriteString("123")
	var a uint16
	a = 1
	s := "12345我"
	fmt.Println([]byte(s))

	writer.WriteString(s)
	binary.Write(writer, binary.LittleEndian, a)
	//binary.Write(writer, binary.LittleEndian, s)

	a = 0
	s = ""
	binary.Read(writer, binary.LittleEndian, &a)
	//binary.Read(writer, binary.LittleEndian, &s)
	fmt.Println(a, s, len(b), writer.Bytes())
}
