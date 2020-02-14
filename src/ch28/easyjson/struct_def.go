/*
 * @Description:
 * @Author: leo
 * @Date: 2020-02-14 17:08:48
 * @LastEditors: leo
 * @LastEditTime: 2020-02-14 17:09:50
 */
package easyjson

type BasicInfo struct {
	Name string
	Age  int
}
type JobInfo struct {
	Skills []string
}

type Employee struct {
	BasicInfo BasicInfo
	JobInfo   JobInfo
}
