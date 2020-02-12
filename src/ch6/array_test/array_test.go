/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-12 15:08:25
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 15:20:23
 */
package array_test

import (
	"fmt"
	"testing"
)

// array 初始化
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5} // 可以通过 ... 来自动得到数组的长度
	arr1[1] = 5
	fmt.Println(arr[1], arr[2])
	fmt.Println(arr1, arr3)
}

// array 循环
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	// for 遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println("for==========", arr3[i])
	}
	// for...range 遍历
	for _, e := range arr3 {
		fmt.Println("for======range", e)
	}
}

// array 的截取
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	arr3_sec := arr3[:]
	arr3_sec1 := arr3[1:]
	arr3_sec2 := arr3[:3]
	arr3_sec3 := arr3[1:len(arr3)]
	fmt.Println("arr3_sec=========", arr3_sec)
	fmt.Println("arr3_sec1=========", arr3_sec1)
	fmt.Println("arr3_sec2=========", arr3_sec2)
	fmt.Println("arr3_sec3=========", arr3_sec3)
}
