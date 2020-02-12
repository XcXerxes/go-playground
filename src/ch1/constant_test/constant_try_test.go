/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-12 13:59:54
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 14:07:48
 */
package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T) {
	fmt.Println(Monday, Tuesday, Wednesday)
}
