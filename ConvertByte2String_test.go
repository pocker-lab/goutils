package Goutils

import (
	"fmt"
	"testing"
)

func TestConvertByte2String(t *testing.T) {
	gb18030 := []byte{192, 180, 215, 212, 32, 49, 52, 46, 49, 49, 57, 46, 49, 48, 52, 46, 49, 56, 57, 32, 181, 196}
	str := ConvertByte2String(gb18030)
	fmt.Println(str)
}
