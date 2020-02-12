/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-12 18:46:50
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 19:09:34
 */
package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	// e := Employee{"1", "Bob", 20}
	// e1 := Employee{Name: "Mike", Age: 30}
	// e2 := new(Employee) // 注意这里返回的引用/指针，相当于 e := &Employee{}
	// e2.Id = "2"
	// e2.Age = 22
	// e2.Name = "Rose"
	// fmt.Println(e)
	// fmt.Println(e1)
	// fmt.Println(e1.Id)
	// fmt.Println(e2)
	// fmt.Println("e is ", e)
	// fmt.Println("e2 is ", e2)
}

// 结构体方法定义
func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// func (e Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	// fmt.Println(e.String())
}
