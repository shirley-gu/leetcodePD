package main

import "fmt"

func main(){
	//solution2-双指针
	n:=[]int{1,3,5,4,7,6,4,2}
	max:=maxArea2(n)
	fmt.Printf("最大面积是：%v", max)
}