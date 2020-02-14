package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		// 这种方式 值会被共享，
		// go func() {
		// 	fmt.Println(i)
		// }()
	}
	time.Sleep(time.Millisecond * 50)
}
