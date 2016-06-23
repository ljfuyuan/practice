/*=============================================================================
#     FileName: ShellSort.go
#         Desc:
#       Author: ljfuyuan
#        Email: ljfuyuan@qq.com
#       Create: 2016-06-23 11:12:06
#   LastChange: 2016-06-23 14:13:45
#      History:
=============================================================================*/

package algorithm

//use Shell's original sequence
func ShellSort_1(data []int) {
	dataLen := len(data)
	for i := dataLen >> 1; i > 0; i >>= 1 {
		for j := i; j < dataLen; j++ {
			tempV := data[j]
			k := 0
			for k = j - i; k >= 0 && tempV < data[k]; k -= i {
				data[k+i] = data[k]
			}
			data[k+i] = tempV
		}
	}
}

//use Knuth's increments
func ShellSort_2(data []int) {
	dataLen := len(data)

	h := 1
	for {
		if h >= dataLen {
			break
		}
		h = h*3 + 1
	}

	for h > 0 {
		for j := h; j < dataLen; j++ {
			tempV := data[j]
			k := 0
			for k = j - h; k >= 0 && tempV < data[k]; k -= h {
				data[k+h] = data[k]
			}
			data[k+h] = tempV
		}
		h = (h - 1) / 3
	}

}
