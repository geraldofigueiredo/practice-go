package insertionsort

type Node struct {
	Val  int
	Next *Node
}

func SortedInsert(newNode *Node, sortedList *Node) *Node {
	// if sortedList == nil {
	// 	sortedList = newNode
	// 	return sortedList
	// }
	if sortedList == nil || sortedList.Val >= newNode.Val {
		newNode.Next = sortedList
		sortedList = newNode
	} else {

		currNode := sortedList
		for currNode.Next != nil && currNode.Next.Val < newNode.Val {
			currNode = currNode.Next
		}

		newNode.Next = currNode.Next
		currNode.Next = newNode
	}

	return sortedList
}

func InsertionSortList(head *Node) *Node {
	var sortedList *Node = nil
	var currNode *Node = &Node{
		Val:  head.Val,
		Next: head.Next,
	}

	for currNode != nil {
		next := currNode.Next
		sortedList = SortedInsert(currNode, sortedList)
		currNode = next
	}

	return sortedList
}
