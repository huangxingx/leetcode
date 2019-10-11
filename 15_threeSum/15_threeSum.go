package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	resultList := make([][]int, 0)

	sort.Ints(nums)

	for k := 0; k < len(nums)-1; k++ {
		if nums[k] > 0 {break}

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
