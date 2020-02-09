/*
指针

*/
package main
import (
	"fmt"
)

func main()  {
	i, j := 42, 2701

	p := &i // 指向 i
	fmt.Println("===================p", *p) // 通过指针读取 i 的值

	*p = 21 // 通过指针设置 i 的值
	fmt.Println("===================i", i) // 查看 i 的值

	p = &j
	*p = *p / 37 // 通过指针进行运算
	fmt.Println("==================j", j)
}
