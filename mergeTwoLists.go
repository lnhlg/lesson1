package lesson1

//ListNode： 链表节点
/*field
Val: 节点值
Next: 后驱节点
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//合并两个有序链表
/*parameter
list1: 有序链表1
list2: 有序链表2
return: 新的有序链表
*/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	//如果有一个链表为空则返回另一链表，两个都为空则返回Nil
	if list1 == nil {

		return list2
	}

	if list2 == nil {

		return list1
	}

	newList := &ListNode{
		Val: -100,
	}
	node := newList

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {

			node.Next = list1
			list1 = list1.Next
		} else {

			node.Next = list2
			list2 = list2.Next
		}

		node = node.Next
	}

	if list1 == nil {

		node.Next = list2
	} else {

		node.Next = list1
	}

	return newList.Next
}
