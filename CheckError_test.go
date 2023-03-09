package Goutils

import (
	"errors"
	"fmt"
	"testing"
)

func TestCheckError(t *testing.T) {
	divide := func(x, y int) (int, error) {
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	}
	result, err := divide(10, 0)
	CheckError(err)
	fmt.Println(result)
}
