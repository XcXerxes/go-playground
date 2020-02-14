/*
 * @Description: 单元测试
 * @Author: leo
 * @Date: 2020-02-14 14:18:39
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-14 14:33:03
 */
package unit

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d", inputs[i], expected[i], ret)
		} else {
			t.Log("pass---------")
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}
func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}
