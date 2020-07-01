/**
 * Auth :   liubo
 * Date :   2020/4/3 18:49
 * Comment:
 */

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	var b = bytes.NewBuffer([]byte{})
	binary.Write(b, binary.BigEndian, int32(1))
	binary.Write(b, binary.BigEndian, int16(2))
	binary.Write(b, binary.BigEndian, float32(3.33))

	var b2 = bytes.NewBuffer(b.Bytes())
	var va int32
	var vb int16
	var vc float32
	binary.Read(b2, binary.BigEndian, &va)
	binary.Read(b2, binary.BigEndian, &vb)
	binary.Read(b2, binary.BigEndian, &vc)

	fmt.Println(va, vb, vc)
}
