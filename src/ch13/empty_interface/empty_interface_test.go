package empty_interface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Interge", i)
	// 	return
	// }
	// if i, ok := p.(string); ok {
	// 	fmt.Println("string", i)
	// 	return
	// }
	// fmt.Println("Unkonw Type")
	// 简化写法
	switch v := p.(type) {
	case int:
		fmt.Println("Interge", v)
	case string:
		fmt.Println("Interge", v)
	default:
		fmt.Println("Unkonw Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("Hello world")
}
