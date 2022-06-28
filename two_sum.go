// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	var res []int
	if len(nums) == 2 {
		res = []int{0, 1}
	} else {
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					res = []int{i, j}
				}
			}
		}
	}
	return res
}