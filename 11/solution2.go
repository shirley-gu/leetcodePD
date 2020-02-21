/*给定 n 个非负整数 a1，a2，...，an，
每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
*/

//双指针法
package main

func maxArea2(height []int) int {
	i, j := 0, len(height)-1
	max := 0

	for i <= j {
		if height[i] > height[j] {
			temp := height[j] * (j - i)
			j--
			if temp >= max {
				max = temp
			}
		} else {
			temp := height[i] * (j - i)
			i++
			if temp >= max {
				max = temp
			}
		}
	}
	return max
}

/*
解析：
https://leetcode-cn.com/problems/container-with-most-water/solution/container-with-most-water-shuang-zhi-zhen-fa-yi-do/

时间复杂度 O(N)O(N)，双指针遍历一次底边宽度 NN 。
空间复杂度 O(1)O(1)，指针使用常数额外空间。
*/