package main

// import "fmt"

// func twoSum2(nums map[int]int, target int) []int {

// 	for k1, v1 := range nums {
// 		for k2, v2 := range nums {
// 			//fmt.Println("k1=", k1)
// 			if v1 + v2 == target && k1 < k2{
// 				fmt.Printf("\t 两个数是%v和%v, 下标分别是%v,%v\n", v1, v2, k1, k2)
// 			}
// 		}
// 	}
// 	return nil
// }
func main(){
	target := 6
	nums := make(map[int]int, 4)
	nums[0] = 2
	nums[1] = 4
	nums[2] = 5
	nums[3] = 7
	twoSum2(nums, target)
}