package Goutils

import "log"

// CheckError 建立一个函数来处理错误返回值，避免写多次重复代码
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
	return
}
