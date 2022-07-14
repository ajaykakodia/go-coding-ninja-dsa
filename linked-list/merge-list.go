package main

func mergeSort(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	midNode := midPointOfLinkedList(head)
	ll1 := head
	ll2 := midNode.next
	midNode.next = nil
	sortedll1 := mergeSort(ll1)
	sortedll2 := mergeSort(ll2)
	sortedlinkedList := mergeSortedLinkedList(sortedll1, sortedll2)
	return sortedlinkedList
}

func mergeSortedLinkedList(link1, link2 *node) *node {
	var link3, currentNode, head *node

	for link1 != nil && link2 != nil {
		if link1.data < link2.data {
			currentNode = link1
			link1 = link1.next
			if link3 != nil {
				link3.next = currentNode
				link3 = currentNode
			} else {
				head = currentNode
				link3 = currentNode
			}

			continue
		} else {
			currentNode = link2
			link2 = link2.next
			if link3 != nil {
				link3.next = currentNode
				link3 = currentNode
			} else {
				head = currentNode
				link3 = currentNode
			}
		}

	}

	if link1 != nil {
		link3.next = link1
	}

	if link2 != nil {
		link3.next = link2
	}

	return head
}
