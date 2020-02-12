/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-12 14:46:08
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 14:47:19
 */
package loop

import (
	"fmt"
	"testing"
)

// while 循环测试
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}
}
