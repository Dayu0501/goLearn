package cn
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 3000 
// -105 <= nums[i] <= 105 
// 
// Related Topics 数组 双指针 
// 👍 3331 👎 0



//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		l, r, n1 := i + 1, len(nums) - 1, nums[i]
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1 + n2 + n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}

				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1 + n2 + n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
