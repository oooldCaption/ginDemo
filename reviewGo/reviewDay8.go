package reviewGo

func RunDay8() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumber(l1, l2 *ListNode) *ListNode {

	// 创建一个新的 链表节点. 作为结果链表的头结点
	head := &ListNode{Val: 0}

	// n1 n2 用来存储 l1,l2当前节点的值; carry 用来存储进位; current 用来跟踪结果链表最后一个值
	n1, n2, carry, current := 0, 0, 0, head

	//这个循环会持续进行，直到两个链表都结束，且没有进位。
	for l1 != nil || l2 != nil || carry != 0 {
		//首先检查 l1 和 l2 当前节点是否为空，如果为空，我们用 0 作为该节点的值。
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		//创建一个新节点，其中的值是 (n1 + n2 + carry) % 10。这实际上是模拟了两数相加并加上进位的操作。
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		// 新创建的节点添加到结果链表的末尾。
		current = current.Next
		//计算新的进位值：(n1 + n2 + carry) / 10
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
