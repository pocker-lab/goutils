package Goutils

import (
	"fmt"
	"os"
	"time"
)

// WriteFile str: 文件内容；打开或创建一个文件，指定写入模式为追加（os.O_APPEND）和创建（os.O_CREATE），权限为 0666
func WriteFile(str, sname string) {
	name := fmt.Sprintf("%v.%v", time.Now().Format("2006-01-02-15-04-05"), sname)
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err) // 如果出错，打印错误并退出程序
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file) // 延迟关闭文件

	_, err = file.WriteString(str) // 将字符串写入到文件中，并返回写入的字节数和错误值
	if err != nil {
		fmt.Println(err) // 如果出错，打印错误并退出程序
		return
	}
}
