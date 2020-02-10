/*
	方法：
		方法是一类带特殊 接收者 参数的函数

		一般 方法接收者 在 func 关键字和方法名之间
		func (v Vertex) Abs() // (v Vertext) 这个就是一个 名为 v, 类型为 Vertex 的接收者
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
