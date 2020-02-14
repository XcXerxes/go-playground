/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-14 13:19:34
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-14 13:25:39
 */
package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Println("obj==========", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
