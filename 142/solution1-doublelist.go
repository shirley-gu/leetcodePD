/*141. 环形链表2
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 
如果 pos 是 -1，则在该链表中没有环。
说明：不允许修改给定的链表。
输入：head = [3,2,0,-4]
输出：tail connects to node index 1
*/

//solution1-双指针
package main

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
 
func detectCycle(head *ListNode) *ListNode {
    // 步骤一：使用快慢指针判断链表是否有环
    slow, fast := head, head
    hasCycle := false
    for slow !=nil && fast!=nil && fast.Next!=nil{
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast{
            hasCycle = true
            break
        }
    }
    // 步骤二：若有环，找到入环开始的节点
    if (hasCycle) {
        for slow != head{
            slow, head = slow.Next, head.Next
        }
        return slow
    } else {
        return nil
    }
}

