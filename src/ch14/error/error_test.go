package error_test

import (
	"errors"
	"fmt"
	"testing"
)

// 区分错误类型
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not large than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetError(t *testing.T) {
	if v, err := GetFibonacci(5); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		fmt.Println(v)
	}
}
