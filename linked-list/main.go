package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("message from linked list")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	line := in.Text()
	strarr := strings.Split(line, " ")
	var head, currentNode *node
	var arr []int
	for _, v := range strarr {
		val, err := strconv.Atoi(v)
		if err == nil {
			arr = append(arr, val)

			newN := newNode(val)

			if head == nil {
				head = newN
				currentNode = newN
				continue
			}
			currentNode.next = newN
			currentNode = newN
		}
	}
	fmt.Println(arr)
	printLinkedList(head)
	sortedLinkedList := mergeSort(head)
	printLinkedList(sortedLinkedList)
	l := length(head)
	fmt.Println("length of linked list :", l)
	midPoint := midPointOfLinkedList(head)
	fmt.Println("midPoint of linked list :", midPoint.data)
	head = deleteNodeFromIthPlace(head, 2)
	printLinkedList(head)
	head = deleteNodeFromIthPlace(head, 0)

	printLinkedList(head)
	head = insertAtIthPosition(head, 0, 9)
	printLinkedList(head)
	head = insertAtIthPRec(head, 2, 11)
	printLinkedList(head)

	deleteNodeFromIthPlaceRec(head, 2)
	printLinkedList(head)
	l = length(head)
	fmt.Println("length of linked list :", l)
	head = reverseLinkedList(head)
	printLinkedList(head)
	head = reverseLinkedList(head)
	printLinkedList(head)
}
