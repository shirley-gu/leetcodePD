/*13. 罗马数字转整数
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

/*
规律：
II=2=1+1
XII=
*/

//violence解法
package main

func romanToInt(s string) int {
	result:=0
	num:=0
	for i:=0; i<len(s); i++ {
		switch s[i] {
			case 'I':
				num = 1
			case 'V':
				num = 5
			case 'X':
				num = 10
			case 'L':
				num = 50
			case 'C':
				num = 100
			case 'D':
				num = 500
			case 'M':
				num = 1000
		}
		if s[i]=='I'&& (s[i+1]=='V'||s[i+1]=='X') ||s[i]=='X'&& (s[i+1]=='L'||s[i+1]=='C')||s[i]=='C'&& (s[i+1]=='D'||s[i+1]=='M'){
			num=num*(-1)
		}
		result= result + num
	}
	return result
}


