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

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (g *JavaProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPloymorphism(t *testing.T) {
	// goProg := &GoProgrammer{} 两种声明方式
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
