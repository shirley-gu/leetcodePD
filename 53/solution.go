/*53. 最大子序和
给定一个整数数组 nums ，
找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/
package pd


func maxSubArray(nums []int) int {
    max:=nums[0]
    sum:=0
    for _,v := range nums{
        if sum > 0{
            sum += v
        }else{
            sum = v
        }

        if sum > max{
            max = sum
        }
    }
    return max
}
