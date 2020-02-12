package condition_test

import (
	"fmt"
	"testing"
)

// 测试 if 支持 在判断结构中 进行两段的形式写法
func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		fmt.Println("1==1")
	}
	// 在函数中的用法
	// if a, err := test(); err == nil {
	// 	// ...
	// } else {
	// 	// ...
	// }
}

// switch 条件测试

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			fmt.Println("switch================Even")
		case 1, 3:
			fmt.Println("switch================Odd")
		default:
			fmt.Println("switch================it is not 0-3")
		}
	}
}

// switch
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			fmt.Println("condition================Even")
		case i%2 == 1:
			fmt.Println("condition================Odd")
		default:
			fmt.Println("condition================it is not 0-3")
		}
	}
}
