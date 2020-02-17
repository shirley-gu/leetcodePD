/*给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那 两个 整数，
并返回他们的数组下标。
*/

//map解法
package main
import "fmt"

//func twoSum(nums map[int]int, target int) []int {
func main() {
	target := 6
	nums := make(map[int]int, 4)
	nums[0] = 2
	nums[1] = 4
	nums[2] = 5
	nums[3] = 7
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			//fmt.Println("k1=", k1)
			if v1 + v2 == target && k1 < k2{
				fmt.Printf("\t 两个数是%v和%v, 下标分别是%v,%v\n", v1, v2, k1, k2)
			}
		}
		//fmt.Println()
	}
}