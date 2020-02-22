/*66. 加一
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
package main

import "fmt"

func plusOne(digits []int) []int {
	for i:=len(digits)-1; i>=0; i--{
		if digits[i]==9{
			digits[i]=0
		}else{
			digits[i]+=1
			break
		}
	}
	return digits
}

func main(){
	n:=[]int{2,1,3,1,2,9,9}
	plusOne(n)
	fmt.Println(n)
}