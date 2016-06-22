/*=============================================================================
#     FileName: InsertSort_test.go
#         Desc:
#       Author: ljfuyuan
#        Email: ljfuyuan@qq.com
#       Create: 2016-06-20 18:19:27
#   LastChange: 2016-06-22 19:55:42
#      History:
=============================================================================*/
package algorithm

import (
	"testing"
)

func insertSort_IsEqual(data []int) (status bool) {
	status = true
	var correctData = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for k, v := range data {
		if v != correctData[k] {
			status = false
			break
		}
	}
	return status
}

func Test_InsertSort_1(t *testing.T) {
	var data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	data = InsertSort_1(data)
	if !insertSort_IsEqual(data) {
		t.Error("InsertSort_1 is not equal to correct answer", data)
	}
}

func Test_InsertSort_2(t *testing.T) {
	var data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	InsertSort_2(data)
	if !insertSort_IsEqual(data) {
		t.Error("InsertSort_2 is not equal to correct answer", data)
	}
}
