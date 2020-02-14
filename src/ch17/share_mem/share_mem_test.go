package share_mem

import (
	"fmt"
	"time"
	"testing"
	"sync"
)

func TestCOunter(t *testing.T)  {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func ()  {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("counter=", counter)
}
// 协程锁
func TestCOunterThreadSafe(t *testing.T)  {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func ()  {
			defer func ()  {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("counter=", counter)
}

// 等待 waitGroup
func TestCOunterWaitGroup(t *testing.T)  {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		// 增加等待的量
		wg.Add(1)
		go func ()  {
			defer func ()  {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			// 等待完成
			wg.Done()
		}()
	}
	// 一直等待到 协程事务完成
	wg.Wait()
	fmt.Println("counter=", counter)
}
