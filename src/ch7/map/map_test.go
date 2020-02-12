/*
 * @Description: map类型
 * @Author: leo
 * @Date: 2020-02-12 15:55:19
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 16:06:30
 */
package map_test

import (
	"fmt"
	"testing"
)

// map 初始化
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	fmt.Println(m1[2])
	fmt.Println("len m1", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	fmt.Println("len m2", len(m2))

	m3 := make(map[int]int, 10)
	fmt.Println("len m3", len(m3))
}

// map 判断key 是否存在 或者是0值
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	fmt.Println(m1[1])
	m1[2] = 0
	fmt.Println(m1[2])

	m1[3] = 1
	if v, ok := m1[3]; ok {
		fmt.Println("key 3 is existing", v)
	} else {
		fmt.Println("key 3 is not existing")
	}
}

// map 循环
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
