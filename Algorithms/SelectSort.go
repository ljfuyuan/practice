/*=============================================================================
#     FileName: SelectSort.go
#         Desc:
#       Author: ljfuyuan
#        Email: ljfuyuan@qq.com
#       Create: 2016-06-23 20:08:32
#   LastChange: 2016-06-23 20:23:51
#      History:
=============================================================================*/
package algorithm

func SelectSort(data []int) {
	dataLen := len(data)
	for i:=0;i<dataLen;i++{
		miniIndex := i
		j := i + 1
		for j<dataLen{
			if data[j] < data[miniIndex] {
				miniIndex = j
			}
			j++
		}
		data[i],data[miniIndex] = data[miniIndex],data[i]
	}
}
