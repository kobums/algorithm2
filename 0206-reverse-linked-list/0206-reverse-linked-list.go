/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var link *ListNode

    for head != nil {
        temp := head.Next

        head.Next = link

        link = head

        head = temp
    }


    return link
}