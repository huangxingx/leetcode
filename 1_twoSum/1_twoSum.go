package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}

	for j := 0; j < len(nums); j++ {
		num := nums[j]
		result := target - num
		_, ok := numMap[result]
		if ok && numMap[result] != j {
			return []int{j, numMap[result]}
		}
	}

	return []int{0, 0}

}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
