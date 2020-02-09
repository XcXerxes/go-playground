/*
	一个结构体 就是一组字段
	结构体的字段用 点(.) 来访问
*/
package main
import (
	"fmt"
)
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("====================v", v.X)
	v.X = 4
	fmt.Println("====================new V", v.X)
	
	v1 := Vertex{X: 1} // 隐式Y 被赋予 0
	fmt.Println("====================V1", v1.X, v1.Y)
	fmt.Println(Vertex{1, 2})
}