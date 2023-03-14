package Goutils

import (
	"fmt"
	"time"
)

// Timing 传入当前时间`time.Now()`，用于打印运行时间
func Timing(times time.Time) {

	time.Sleep(time.Second)
	elapsed := time.Since(times)
	hours := int(elapsed.Hours())
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60
	fmt.Printf("\r运行时长: %02d:%02d:%02d", hours, minutes, seconds)
}
