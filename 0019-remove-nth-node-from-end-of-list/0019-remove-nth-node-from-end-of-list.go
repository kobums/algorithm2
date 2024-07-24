/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    d, end := getDaddy(head, n)

    if end {
        return head.Next
    }

    d.Next = d.Next.Next

    return head
}

func getDaddy(head *ListNode, n int) (daddy *ListNode, end bool) {
    daddy = head

    for head != nil {
        if n < 0 {
            daddy = daddy.Next
        }
        n--
        head = head.Next
    }

    end = n == 0
    return
}