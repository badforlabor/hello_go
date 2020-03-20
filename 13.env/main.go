/**
 * Auth :   liubo
 * Date :   2019/10/15 9:46
 * Comment: 测试环境变量
 */

package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

func main() {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment", registry.ALL_ACCESS)
	if err == nil {
		v, _, err := key.GetStringValue("Path")
		if err == nil {
			fmt.Printf("系统环境变量：%s\n", v)
		}
		key.Close()
	}
	key, err = registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	if err == nil {
		v, _, err := key.GetStringValue("Path")
		if err == nil {
			fmt.Printf("用户环境变量：%s", v)
		}
		key.SetStringValue("aaa", "bbb")
		key.DeleteValue("aaa")

		key.Close()
	}
}