/*给定一个包含 n 个整数的数组 nums，
判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
*/

package pd

import "Sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	addsum := [][]int{}
    for i :=0; i < len(nums); i++ {
		for j :=i+i; j < len(nums); j++ {
			for k :=j+1; k < len(nums); k++ {
				if nums[i] + nums[j] + nums[k]== 0 {
					addsum = append(addsum, []int{nums[i], nums[j], nums[k]})
				}
			}		
		}
	}
	return addsum
}

/*未去重
func threeSum(nums []int) [][]int {
    for i :=0; i < len(nums); i++ {
		for j :=i+i; j < len(nums); j++ {
			for k :=j+1; k < len(nums); k++ {
				if nums[i] + nums[j] + nums[k]== 0 {
					addsum = append(addsum, []int{nums[i], nums[j], nums[k]})
				}
			}		
		}
	}
	return addsum
}
*/