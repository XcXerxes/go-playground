/*
	map 的语法与结构体相似，不过必须有键名
*/
package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

// 定义一个map
var m = map[string]Vertex{
	"Bell Labs": Vertex{40.0, -40.1},
	"Google":    Vertex{37.22, -100.33},
}

// 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
var m1 = map[string]Vertex{
	"Bell Labs": {40.0, -40.1},
	"Google":    {37.22, -100.33},
}

func main() {
	fmt.Println(m)
	fmt.Println(m1)
	// 修改或者插入 元素
	m2 := make(map[string]int)
	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])
	// 删除元素
	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])
	// 检测元素是否存在
	m2["Answer"] = 20
	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
