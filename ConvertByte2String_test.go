package Goutils

import (
	"fmt"
	"testing"
)

func TestConvertByte2String(t *testing.T) {
	gb18030 := []byte{192, 180, 215, 212, 32, 49, 52, 46, 49, 49, 57, 46, 49, 48, 52, 46, 49, 56, 57, 32, 181, 196, 187, 216, 184, 180, 58, 32, 215, 214, 189, 218, 61, 49, 48, 48, 32, 202, 177, 188, 228, 61, 51, 48, 109, 115, 32, 84, 84, 76, 61, 53, 52, 13, 10}
	str := ConvertByte2String(gb18030)
	fmt.Println(str)
}
