// 类型转换

package main
import (
	"fmt"
	"math"
)

func main() {
	// 数值的转换
	x, y := 3, 4
	// 将 int 转为 float64
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// 浮点类型 转为 uint
	z := uint(f)

	// 如果 声明变量 不指定类型 由右值 推导得出
	i := 42 // int
	f := 3.142 // float64
	fmt.Println(x, y, z)
}