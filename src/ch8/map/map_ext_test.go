/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-12 16:08:44
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 16:18:30
 */
package map_test

import (
	"fmt"
	"testing"
)

// map 的值为 func
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](2), m[2](2), m[3](2))
}

// map 实现 Set

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true

	if n := 3; mySet[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}
	mySet[3] = false
	fmt.Println("=======len", len(mySet))
	delete(mySet, 1)
	if n := 1; mySet[n] {
		fmt.Printf("%d is existing", n)
	} else {
		fmt.Printf("%d is not existing", n)
	}
}
