package concepts

// ListNode is a node
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverses a linked list
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return temp
}

// GetRow returns nth row in pascals triangle
func GetRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	oldRow := GetRow(rowIndex - 1)
	nthRow := make([]int, rowIndex+1)
	nthRow[0] = 1
	for i := 1; i < rowIndex; i++ {
		nthRow[i] = oldRow[i] + oldRow[i-1]
	}
	nthRow[rowIndex] = 1
	return nthRow
}
