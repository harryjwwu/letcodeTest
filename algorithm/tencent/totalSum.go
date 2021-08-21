package tencent

import "fmt"

/*
一. 基础难度，一星：基础

给定一个数组，找出 和为目标值的两个整数，返回这两个整数的下标。例如给定数组[3, 1, 2 ,6] ，目标值为 9； 3+6 =9， 故返回下标[0, 3]
补充以下函数实现
vector<int> twoSum(vector<int>& nums, int sum) {
}
空间复杂度：O(n)
时间复杂度：O(n)

 */

func TotalSum(nums []int32 , sum int32) []int {
	var valueMap map[int32]int
	valueMap = make(map[int32]int)
	// value - index
	for i := 0; i < len(nums); i++ {
		valueMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if iIndex, ok := valueMap[sum - nums[i]]; ok {
			fmt.Printf("index:%d, %d", iIndex, i)
			return []int{iIndex, i}
		}
	}
	return nil
}
