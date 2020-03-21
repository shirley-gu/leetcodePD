//solution2-map哈希
//参考看
package main

func detectCycle2(head *ListNode) *ListNode {
    nodeExist := make(map[*ListNode]int)
    cur := head
    for cur != nil {
        if _, exist := nodeExist[cur]; exist {
            return cur
        }
        nodeExist[cur] = 1
        cur = cur.Next
    }
    return nil    
}