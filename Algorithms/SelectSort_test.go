/*=============================================================================
#     FileName: SelectSort_test.go
#         Desc:
#       Author: ljfuyuan
#        Email: ljfuyuan@qq.com
#       Create: 2016-06-23 20:20:06
#   LastChange: 2016-06-23 20:22:11
#      History:
=============================================================================*/
package algorithm

import (
	"testing"
)

func Test_SelectSort_1(t *testing.T) {
	var data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	SelectSort(data)
	if !isEqual(data) {
		t.Error("SelectSort is not equal to correct answer", data)
	}
}

