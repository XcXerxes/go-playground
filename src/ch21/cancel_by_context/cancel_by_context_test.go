/*
 * @Description: 123
 * @Author: leo
 * @Date: 2020-02-14 12:57:03
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-14 13:16:54
 */
package cancel_by_context_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	// 判断任务是否被取消
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 50)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	// 取消任务
	cancel()
	time.Sleep(time.Second * 1)
}
