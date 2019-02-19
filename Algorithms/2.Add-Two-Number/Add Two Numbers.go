package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	head := &ListNode{Val: (l1.Val + l2.Val) % 10, Next: nil}
	last := head
	carry := (l1.Val + l2.Val) / 10
	l1 = l1.Next
	l2 = l2.Next
	if l1 == nil && l2 != nil {
		for l2.Next != nil {
			temp := &ListNode{Val: (l2.Val + carry) % 10, Next: nil}
			carry = (l2.Val + carry) / 10
			last.Next = temp
			last = temp
			l2 = l2.Next
		}
		temp := &ListNode{Val: (l2.Val + carry) % 10, Next: nil}
		carry = (l2.Val + carry) / 10
		last.Next = temp
		last = temp
		if carry > 0 {
			temp := &ListNode{Val: carry, Next: nil}
			last.Next = temp
		}
		return head
	} else if l1 != nil && l2 == nil {
		for l1.Next != nil {
			temp := &ListNode{Val: (l1.Val + carry) % 10, Next: nil}
			carry = (l1.Val + carry) / 10
			last.Next = temp
			last = temp
			l1 = l1.Next
		}
		temp := &ListNode{Val: (l1.Val + carry) % 10, Next: nil}
		carry = (l1.Val + carry) / 10
		last.Next = temp
		last = temp
		if carry > 0 {
			temp := &ListNode{Val: carry, Next: nil}
			last.Next = temp
		}
		return head
	} else if l1 == nil && l2 == nil && carry > 0 {
		temp := &ListNode{Val: carry, Next: nil}
		last.Next = temp
		return head
	} else if l1 == nil && l2 == nil && carry <= 0 {
		return head
	} else {
		for l1.Next != nil && l2.Next != nil {
			temp := &ListNode{Val: (l1.Val + l2.Val + carry) % 10, Next: nil}
			carry = (l1.Val + l2.Val + carry) / 10
			last.Next = temp
			last = temp
			l1 = l1.Next
			l2 = l2.Next
		}
		if l1.Next == nil && l2.Next != nil {
			temp := &ListNode{Val: (l1.Val + l2.Val + carry) % 10, Next: nil}
			carry = (l1.Val + l2.Val + carry) / 10
			last.Next = temp
			last = temp
			l2 = l2.Next
			for l2.Next != nil {
				temp := &ListNode{Val: (l2.Val + carry) % 10, Next: nil}
				carry = (l2.Val + carry) / 10
				last.Next = temp
				last = temp
				l2 = l2.Next
			}
			temp = &ListNode{Val: (l2.Val + carry) % 10, Next: nil}
			carry = (l2.Val + carry) / 10
			last.Next = temp
			last = temp
			if carry > 0 {
				temp := &ListNode{Val: carry, Next: nil}
				last.Next = temp
			}
			return head
		} else if l2.Next == nil && l1.Next != nil {
			temp := &ListNode{Val: (l1.Val + l2.Val + carry) % 10, Next: nil}
			carry = (l1.Val + l2.Val + carry) / 10
			last.Next = temp
			last = temp
			l1 = l1.Next
			for l1.Next != nil {
				temp := &ListNode{Val: (l1.Val + carry) % 10, Next: nil}
				carry = (l1.Val + carry) / 10
				last.Next = temp
				last = temp
				l1 = l1.Next
			}
			temp = &ListNode{Val: (l1.Val + carry) % 10, Next: nil}
			carry = (l1.Val + carry) / 10
			last.Next = temp
			last = temp
			if carry > 0 {
				temp := &ListNode{Val: carry, Next: nil}
				last.Next = temp
			}
			return head
		} else if l1.Next == nil && l2.Next == nil {
			temp := &ListNode{Val: (l1.Val + l2.Val + carry) % 10, Next: nil}
			carry = (l1.Val + l2.Val + carry) / 10
			last.Next = temp
			last = temp
			if carry > 0 {
				temp := &ListNode{Val: carry, Next: nil}
				last.Next = temp
			}
			return head
		}
	}
	return head
}
