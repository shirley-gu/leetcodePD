/*206. 反转链表
反转一个单链表。
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

//递归解法
package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverse(pre,cur *ListNode) *ListNode {
    if cur == nil {
        return pre
    }
    head := reverse(cur, cur.Next)
    cur.Next = pre
    return head
}

func reverseList(head *ListNode) *ListNode {    // 递归方法，原地反转链表
    return reverse(nil, head)
}