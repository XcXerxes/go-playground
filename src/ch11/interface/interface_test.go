/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-12 19:16:16
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 19:20:05
 */
package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.WriteHelloWorld())
}
