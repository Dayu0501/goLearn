package leetcode

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	threeSum(nums)
}

func threeSum(nums []int) [][]int {
	//var res [][]int
	//sort.Ints(nums)
	//
	////println(nums)
	//
	//for i := 0; i < len(nums) - 2; i++ {
	//	if nums[i] > 0 || nums[len(nums) - 1] < 0 {
	//		break
	//	}
	//
	//	if i > 0 && nums[i] == nums[i-1] {
	//		continue
	//	}
	//	l, r := i + 1, len(nums) - 1
	//	for l < r {
	//		sum := nums[i] + nums[l] + nums[r]
	//		if sum < 0 {
	//			for l < r {
	//				l++
	//			}
	//		} else if sum > 0 {
	//			for l < r{
	//				r--
	//			}
	//		} else {
	//			res = append(res, []int{nums[i], nums[l], nums[r]})
	//			for l < r && nums[l] == nums[l+1] {
	//				l++
	//			}
	//			for l < r && nums[r] == nums[r-1] {
	//				r--
	//			}
	//		}
	//	}
	//}
	//return res

	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
