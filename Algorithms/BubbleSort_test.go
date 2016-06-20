/*=============================================================================
#     FileName: BubbleSort_test.go
#         Desc:
#       Author: ljfuyuan
#        Email: ljfuyuan@qq.com
#       Create: 2016-06-20 17:10:55
#   LastChange: 2016-06-20 17:10:57
#      History:
=============================================================================*/
package algorithm

import (
	"testing"
)

func isEqual(data []int) (status bool) {
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

func Test_BubbleSort_1(t *testing.T) {
	var data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort_1(data)
	if !isEqual(data) {
		t.Error("BubbleSort_1 is not equal to correct answer", data)
	}
}

func Test_BubbleSort_2(t *testing.T) {
	var data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort_2(data)
	if !isEqual(data) {
		t.Error("BubbleSort_2 is not equal to correct answer", data)
	}
}

func Test_BubbleSort_3(t *testing.T) {
	var data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort_3(data)
	if !isEqual(data) {
		t.Error("BubbleSort_3 is not equal to correct answer", data)
	}
}
