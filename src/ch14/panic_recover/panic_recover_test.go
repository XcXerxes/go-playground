package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
	}()
	fmt.Println("Start")
	// panic(errors.New("Something wrong!")) // 执行defer
	// os.Exit(-1) // 不执行 defer
}

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("Something wrong!"))
}
