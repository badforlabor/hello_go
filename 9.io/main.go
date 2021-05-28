package main

import (
	"fmt"
	"os"
)

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


