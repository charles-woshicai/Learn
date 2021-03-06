/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

链接：https://leetcode-cn.com/problems/two-sum
*/
package twosum

func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	var anotherMap = make(map[int]int)
	for i, v := range nums {
		another := target - v
		if j, ok := anotherMap[another]; ok {
			return []int{i, j}
		}
		anotherMap[v] = i
	}
	return []int{}
}
