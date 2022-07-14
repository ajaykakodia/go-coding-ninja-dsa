package main

import "fmt"

type linkedList struct {
	head *node
}

type node struct {
	data int
	next *node
}

func newNode(val int) *node {
	return &node{
		data: val,
	}
}

func printLinkedList(head *node) {
	if head == nil {
		return
	}

	for head != nil {
		fmt.Printf("%d -> ", head.data)
		head = head.next
	}
	fmt.Print("null")
	fmt.Println()
}

func length(head *node) int {
	count := 0

	for head != nil {
		count++
		head = head.next
	}
	return count
}

func insertAtIthPosition(head *node, pos, val int) *node {
	newN := newNode(val)
	if pos == 0 {
		newN.next = head
		return newN
	}
	currentNode := head
	for i := 0; i < pos-1 && currentNode != nil; i++ {
		currentNode = currentNode.next
	}
	if currentNode == nil || pos < 0 {
		fmt.Println("Not able to insert at position :", pos)
		return head
	}
	newN.next = currentNode.next
	currentNode.next = newN
	return head
}

func insertAtIthPRec(head *node, pos, val int) *node {

	newN := newNode(val)
	if pos == 0 {
		newN.next = head
		return newN
	}
	if head == nil {
		return head
	}
	head.next = insertAtIthPRec(head.next, pos-1, val)
	return head
}

func deleteNodeFromIthPlace(head *node, pos int) *node {
	if head == nil || pos < 0 {
		return head
	}
	if pos == 0 {
		return head.next
	}

	currentNode := head

	for i := 0; i < pos-1 && currentNode != nil; i++ {
		currentNode = currentNode.next
	}
	if currentNode == nil || currentNode.next == nil {
		return head
	}
	if currentNode.next != nil && currentNode.next.next == nil {
		currentNode.next = nil
		return head
	}
	currentNode.next = currentNode.next.next
	return head
}

func deleteNodeFromIthPlaceRec(head *node, pos int) *node {
	if head == nil {
		return head
	}
	if pos == 0 {
		return head.next
	}
	head.next = deleteNodeFromIthPlaceRec(head.next, pos-1)
	return head
}

func reverseLinkedList(head *node) *node {
	if head != nil && head.next == nil {
		return head
	}

	reversedSmallLinkedList := reverseLinkedList(head.next)

	head.next.next = head
	head.next = nil

	return reversedSmallLinkedList
}

func reverseLinkedListIte(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var temp, currentNode *node
	currentNode = nil
	temp = head
	for temp.next != nil {
		temp = head.next
		head.next = currentNode
		currentNode = head
		head = temp
	}

	return head
}

func midPointOfLinkedList(head *node) *node {

	fastMover, slowMover := head, head
	for fastMover.next != nil && fastMover.next.next != nil {
		slowMover = slowMover.next
		fastMover = fastMover.next.next
	}

	return slowMover
}
