/*
 * @Description: 切片
 * @Author: leo
 * @Date: 2020-02-12 15:20:57
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 15:45:56
 */
package slice_test

import (
	"fmt"
	"testing"
)

// 切片初始化
func TestSliceInit(t *testing.T) {
	var s0 []int // 与数组的区别就可变长度的
	s0 = append(s0, 1)
	fmt.Println(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	fmt.Println(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	fmt.Println(s2[0], s2[1], s2[2])
	fmt.Println(len(s2), cap(s2))
}

// 切片可变长
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
}

// 切片共享 修改任何一个切片的附属切片都会 改变其他的附属切片

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	Q2 := year[3:6]
	fmt.Println(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	fmt.Println(summer, len(summer), cap(summer))

	summer[0] = "Unknow"

	fmt.Println(Q2)
}

// 切片之间不能比较 只能跟 nil 比较
func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3}
	// b := []int{1, 2, 3}

	// 无法比较
	// fmt.Println(a == b)
	fmt.Println(a == nil)
}
