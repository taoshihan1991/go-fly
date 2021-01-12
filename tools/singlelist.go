package tools

// 单链表节点的结构
type ListNode struct {
	val  int
	next *ListNode
}

func NewListNode(x int) *ListNode {
	return &ListNode{
		val: x,
	}
}
func ReverseList(head *ListNode) *ListNode {
	if head.next == nil {
		return head
	}
	last := ReverseList(head.next)
	head.next.next = head
	head.next = nil
	return last
}

var successor *ListNode // 后驱节点
// 将链表的前 n 个节点反转（n <= 链表长度）
func ReverseListN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n + 1 个节点
		successor = head.next
		return head
	}
	// 以 head.next 为起点，需要反转前 n - 1 个节点
	last := ReverseListN(head.next, n-1)
	head.next.next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.next = successor
	return last
}

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return ReverseListN(head, m)
	}
	// 前进到反转的起点触发 base case
	head.next = ReverseBetween(head.next, m-1, n-1)
	return head
}

/** 反转区间 [a, b) 的元素，注意是左闭右开 */
func ReverseSingleList(a *ListNode, b *ListNode) *ListNode {
	var (
		pre, cur, nxt *ListNode
	)
	pre = nil
	cur = a
	nxt = a
	//终止的条件改一下就行了
	for cur != b {
		nxt = cur.next
		// 逐个结点反转
		cur.next = pre
		// 更新指针位置
		pre = cur
		cur = nxt
	}
	// 返回反转后的头结点
	return pre
}
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k 个待反转元素
	var (
		a, b *ListNode
	)
	b = head
	a = head
	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.next
	}
	// 反转前 k 个元素
	newHead := ReverseSingleList(a, b)
	// 递归反转后续链表并连接起来
	a.next = ReverseKGroup(b, k)
	return newHead
}
