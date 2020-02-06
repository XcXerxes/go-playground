package main

import (
	"fmt"
	"math"
)

// if 跟 for 类型 表达式无需 小括号 (), 而大括号 {}是必须的
func sqrt (x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if 也可以在 条件表达式之前 执行一个简单的语句
func pow(x, n, lim float64) float64  {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// for 循环
// 结构体里面是 没有小括号的
func main()  {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	// 初始化语句和后置语句都是可选的
	for ; sum < 100; {
		sum += sum
	}
	// for 是 Go 中的 while
	for sum < 1000 {
		sum += sum
	}
	// defer 语句会将 函数推迟到外层函数返回之后再执行
	defer fmt.Println("defer========", sqrt(5), sqrt(-5))

	for k := 0; k < 10; k++ {
		defer fmt.Println(k)
	}
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	// 无限循环
	// for {
	// 	// 循环内容
	// }
	fmt.Println(sum)
}
