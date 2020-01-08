package main
import "fmt"

// 函数可以接收多个参数
// **** 注意 类型在变量名
// 括号里面是参数 紧跟着的是 返回值
func add(x int, y int) int {
	return x + y
}
// 连续两个或者多个参数类型相同时，除了最后一个要申明类型以外，其他的都可以省略
func add1(x, y int) int {
	return x + y
}
func add2(x, y, z int) int {
	return x + y + z
}

// 函数可以返回任意数量的返回值
// 这里是返回两个字符串
func swap(x, y string) (string, string)  {
	return y, x
}

// 函数的 返回值可被命名，它们被定义为 函数顶部的变量。
// 返回值的名称应当 具有一定意义，它可以作为文档使用
// 没参数的 retrun 语音 直接返回 已命名的返回值
// 直接返回的语句 一般在 短函数中使用，长函数使用的话 会影响代码的可读性
func split(sum int) (x, y int)  {
	x = sum * 4 /9
	y = sum - 4
	return
}
func main() {
	// 打印输出
	fmt.Println(add(10, 12))
	a, b := swap("Hello", "world")
	fmt.Println(a, b)
	x, y := split(17)
	fmt.Println(x, y)
}
