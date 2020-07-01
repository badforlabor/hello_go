package main

import (
	"errors"
	"fmt"
	"os"
)

func Rename(oldpath, newpath string) error {
	var f, e = os.Open(newpath)
	if e != nil || f != nil {
		return errors.New("目录路径非法")
	}
}


func main() {


	var e = os.Rename("E:/dispatcher41", "F:/dispatcher41")
	if e != nil {
		fmt.Println("移动失败!", e.Error())
	}

	return

	testInterface()
	testBuf()

	testDisk()
	listFolder("E:/")

	fmt.Println("test end")
}


