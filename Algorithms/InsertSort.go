/*=============================================================================
#     FileName: InsertSort.go
#         Desc:
#       Author: ljfuyuan
#        Email: ljfuyuan@qq.com
#       Create: 2016-06-20 17:19:00
#   LastChange: 2016-06-22 20:02:40
#      History:
=============================================================================*/
package algorithm

func InsertSort_1(data []int) []int {
	newData := make([]int, 1, cap(data))
	newData[0] = data[0]
	data = data[1:]

	for _, v := range data {
		for m, n := range newData {
			if v <= n {
				newData = append(newData[:m+1], newData[m:]...)
				newData[m] = v
				break
			}
		}
	}
	return newData
}

func InsertSort_2(data []int) {
	for i := 1; i < len(data); i++ {
		tempV := data[i]
		var j int
		for j = i; j > 0 && tempV < data[j-1]; j-- {
			data[j] = data[j-1]
		}
		data[j] = tempV
	}
}
