/*
 * @Description: 反射编程
 * @Author: leo
 * @Date: 2020-02-14 15:18:49
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-14 16:05:38
 */
package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	// map 比较
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}

	fmt.Println("a vs b==========", reflect.DeepEqual(a, b))

	// 切片比较
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	fmt.Println("s1 == s2?", reflect.DeepEqual(s1, s2))
	fmt.Println("s1 == s3?", reflect.DeepEqual(s1, s3))
	// var f int64 = 10
	// fmt.Println(reflect.TypeOf(f), reflect.ValueOf(f))
	// fmt.Println(reflect.ValueOf(f).Type())
}
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	// 按名字获取成员
	fmt.Printf("Name: value(%[1]v), Type(%[1]T)\n", reflect.ValueOf(*e).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Faile")
	} else {
		fmt.Println("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	fmt.Println("Updated Age:", e)
}
