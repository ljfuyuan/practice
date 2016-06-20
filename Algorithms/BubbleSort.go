/*=============================================================================
#     FileName: BubbleSort.go
#         Desc: three kinds of Bubblesort
#       Author: ljfuyuan
#        Email: ljfuyuan@qq.com
#       Create: 2016-06-20 15:13:13
#   LastChange: 2016-06-20 17:10:50
#      History:
=============================================================================*/
package algorithm

func BubbleSort_1(data []int) {
	dataLen := len(data)
	for k := 0; k < dataLen; k++ {
		for i := 1; i < dataLen; i++ {
			if data[i] < data[i-1] {
				data[i], data[i-1] = data[i-1], data[i]
			}
		}
	}
}

func BubbleSort_2(data []int) {
	status := true
	dataLen := len(data)
	for status {
		status = false
		for i := 1; i < dataLen; i++ {
			if data[i] < data[i-1] {
				data[i], data[i-1] = data[i-1], data[i]
				status = true
			}
		}
		dataLen--
	}
}

func BubbleSort_3(data []int) {
	dataLen := len(data)
	lastPos := dataLen
	for lastPos > 0 {
		maxPos := lastPos
		lastPos = 0
		for i := 1; i < maxPos; i++ {
			if data[i] < data[i-1] {
				data[i], data[i-1] = data[i-1], data[i]
				lastPos = i
			}
		}
	}
}
