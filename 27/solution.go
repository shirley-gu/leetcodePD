/*给定一个数组 nums 和一个值 val，
你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/
package main

import "fmt"

func removeElement(nums []int, val int) int {
	p:=0
	for i:=0; i<len(nums); i++{
		if nums[i]== val{
			p++
		}else{
			if p!=0{
				nums[i-p]=nums[i]
			}
		}
	}
	return len(nums)-p
}

func main(){
	n:=[]int{0,1,2,2,3,0,4,2}
	removeElement(n,2)
	fmt.Println(n)
}

/*jenny做法：
func removeElement(nums []int, val int) int {
	j:=0
	for i:=0; i<len(nums); i++{
		if nums[i]!= val{
			nums[j]=nums[i]
			j++
		}
	}
	return j
}
*/