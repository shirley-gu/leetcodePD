/*给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那 两个 整数，
并返回他们的数组下标。
*/
package main

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++{
        for j := i+1; j < len(nums); j++ {
            if nums[i] + nums[j] == target{
                return []int{i,j}
            }
        }
    }
    return nil
}