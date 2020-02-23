package main

import "fmt"

func main(){
	//solution1-violence
	var sum string = "MCMXCIV"
	a:=romanToInt(sum)
	fmt.Println(a)
	//solution2-map
	b:=romanToInt2(sum)
	fmt.Println(b)
}