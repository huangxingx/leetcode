package main

import (
	"fmt"
	"sort"
)

/*
数组排序
两层循环, 外层循环遍历整个数组, 内层循环通过两个游标获取首尾数据
if sum < 0 , 递增左边的游标
if sum > 0 , 递增右边的游标
else append 三个数的数组

注意数据的去重处理
 */

func threeSum(nums []int) [][]int {
	resultList := make([][]int, 0)

	sort.Ints(nums)

	for k := 0; k < len(nums)-1; k++ {
		if nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i, j := k+1, len(nums)-1
		for i < j {
			s := nums[i] + nums[j] + nums[k]
			if s < 0 {
				i += 1
				for nums[i] == nums[i-1] && (i < j) {
					i += 1
				}
			} else if s > 0 {
				j -= 1
				for nums[j] == nums[j+1] && (i < j) {
					j -= 1
				}
			} else {

				resultList = append(resultList, []int{nums[k], nums[i], nums[j]})
				i += 1
				j -= 1

				for nums[j] == nums[j+1] && (i < j) {
					j -= 1
				}
				for nums[i] == nums[i-1] && (i < j) {
					i += 1
				}
			}

		}

	}

	return resultList
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
