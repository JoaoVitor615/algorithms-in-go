package merge_sort

// Node represents a node in the linked list.
type Node struct {
	Value int
	Next  *Node
}

// MergeSort sorts a linked list using the Merge Sort algorithm.
func MergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	// Find the middle of the list to split it.
	slow, fast := head, head
	var prev *Node
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	// Recursively call Merge Sort on the two halves.
	left := MergeSort(head)
	right := MergeSort(slow)

	// Merge the two sorted halves.
	return merge(left, right)
}

// merge combines two sorted linked lists into a single sorted list.
func merge(l1, l2 *Node) *Node {
	dummy := &Node{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}
