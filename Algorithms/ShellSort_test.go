/*=============================================================================
#     FileName: ShellSort_test.go
#         Desc:
#       Author: ljfuyuan
#        Email: ljfuyuan@qq.com
#       Create: 2016-06-23 14:10:28
#   LastChange: 2016-06-23 14:17:43
#      History:
=============================================================================*/

package algorithm

import (
	"testing"
)

func Test_ShellSort_1(t *testing.T) {
	var data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	ShellSort_1(data)
	if !isEqual(data) {
		t.Error("ShellSort_1 is not equal to correct answer", data)
	}
}

func Test_ShellSort_2(t *testing.T) {
	var data = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	ShellSort_2(data)
	if !isEqual(data) {
		t.Error("ShellSort_2 is not equal to correct answer", data)
	}
}
