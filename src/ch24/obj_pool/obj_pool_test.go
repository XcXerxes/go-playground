/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-14 13:57:00
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-14 14:05:40
 */
package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)

	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("Done")
}
