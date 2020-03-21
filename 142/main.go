package main

import "fmt"

func main(){
	list := &ListNode{}
	newlist := detectCycle(list)
	fmt.Println(newlist)
}