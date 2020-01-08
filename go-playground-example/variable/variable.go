package main
import "fmt"

// var 语句用于 声明一个变量列表
// 跟函数参数列表一样，类型在最后
var c, python, java bool

// 变量可以包含 初始值，每个变量对应一个
// 有初始值时，可以省略 类型，变量会从初始值中 获取类型
var sum, calc int = 1, 2

// 只有在以 关键字开始的语句才能使用（var, func 等等） 短变量声明
// b := 3 -> 报错

// 常量的声明 不能用 := 语法声明 使用 const 关键字
const Pi = 3.14
func main() {
	// 函数内部 声明 var
	var i int
	// 内部函数也同样可以
	var js, php = true, "no!"
	fmt.Println(sum, calc, js, php)
	fmt.Println(i, c, python, java)

	// 在函数中 ， 简洁赋值语句 := 可在类型明确的 地方代替 var 声明
	k := 3

	// 常量声明
	const World = "世界"
	fmt.Println(k)
}