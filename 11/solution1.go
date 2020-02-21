/*给定 n 个非负整数 a1，a2，...，an，
每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
*/

//暴力解法-超出时间限制
package main

import "math"

func maxArea(height []int) int{
	if len(height)<2{
		return -1
	}
	len:=len(height)
	result:=math.Min(float64(height[len-1]),float64(height[0]))*float64(len-1)
	for i:=0; i<len; i++{
		for j:=len-1; j>i; j--{
			temp:= math.Min(float64(height[i]),float64(height[j]))*float64(j-i)
			if temp>result{
				result=temp
			}
		}		
	}
	return int(result)	
}

