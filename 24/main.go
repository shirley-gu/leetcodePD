/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
//Definition for singly-linked list.

package main

type ListNode struct {
	Val int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n2 := head.Next
	head.Next = swapPairs(n2.Next)
	n2.Next = head
	return n2
}
//1->2->3->4
/*
head = 1, n2 = 2
1.next= swap(3)
head = 3, n2 = 4
3.next = swap(4.next)   swap(4.next) return 4.next = nil
3.next = 4.next
4.next = 3
return 4
1->2->nil   nil->4->3
1.next = 4
2.next =1
2->1->4->3
*/
	