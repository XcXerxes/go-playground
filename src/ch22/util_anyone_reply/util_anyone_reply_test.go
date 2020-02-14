/*
 * @Description: 任意任务完成 即可返回
 * @Author: leo
 * @Date: 2020-02-14 13:33:37
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-14 13:43:15
 */
package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// 需要使用 带 容量的 协程 防止 协程泄漏
	// ch := make(chan string) 造成 协程泄漏
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	fmt.Println("Before:", runtime.NumGoroutine())
	fmt.Println(FirstResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("After:", runtime.NumGoroutine())
}
