/*
	1、数组： 类型 [n]T 表示拥有 n 个 T 类型的值的数组
	数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。

	2、切片：数组的大小是固定的，通过切片 可以提供动态大小的，灵活的视角
	类型 []T 表示一个元素类型为 T 的切片
	a[low:gigh]
	切片包含一个上界和一个下界，包含第一个元素，但排除最后一个元素

	3、切片就像数组的引用， 它不存储任何数据， 
	更改切片的元素 会修改 其底层数组中对应的元素, 而且与之共享底层数组的切片都会被修改

	4、切片文法类似于没有长度的数组文法
		数组文法：
		[1]bool{true}
		切片文法：
		[]bool{false}
	5、切片的默认行为
		切片下界的默认值为 0，上界则是改切片的长度。
		a[0:10] == a[:10] == a[0] == a[:]
	6、用make 内置函数 创建 切片
		a := make([]int, 0, 5) // 第一个参数是类型 第二个参数是长度，第三个参数是容量
*/
package main
import (
	"fmt"
)

func main()  {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 表示创建一个 primes数组中从 下标 1 到3 的一个切片
	var s []int = primes[1:4]
	fmt.Println(s)

	//
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("=============names", names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	// 切片文法
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("==================q", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("===================r", r)

	f := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("==================f", f)

	// 切片默认行为
	s1 := []int{2, 3, 5, 7, 11, 13}
	s1 = s1[1:4]
	fmt.Println("==================s1", s1)

	s1 = s1[:2]
	fmt.Println("==================s1", s1)

	s1 = s1[1:]
	fmt.Println("====================s1",s1)
	// 切片的长度和容量 通过 len() 和 cap()来获取
	s2 := []int{2, 3, 5, 7, 11, 13}
	// 截取切片使其长度为 0
	s2 = s2[:0]
	printSlice(s2)
	// 拓展其长度
	s2 = s2[:4]
	printSlice(s2)
	// 舍弃前两个值
	s2 = s2[2:]
	printSlice(s2)
	// nil 切片
	var s3 []int
	printSlice(s3)
	

	// 往切片添加元素
	appendEle()

	// 循环遍历切片
	forEle()
}

// append方法 append(s []T, vs ...T) []T
func appendEle()  {
	var s []int
	// 添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

// for 循环 range 形式遍历切片
func forEle() {
	pow := []int{1, 2, 4, 8, 16, 32, 61, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func printSlice(s []int) {
	fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)
}
