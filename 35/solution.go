/*35. 搜索插入位置
给定一个排序数组和一个目标值，
在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。
*/
package main

import "fmt"

func searchInsert(nums []int, target int) int {
	i:=0
	for i=0; i<len(nums); i++{
		if nums[i]>= target{
			break
		}
	}
	return i
}

func main(){
	n:=[]int{0,1,2,4,5}
	i:=searchInsert(n,3)
	fmt.Println(i)
}
