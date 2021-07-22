package main

import (
	"fmt"

	"go.uber.org/zap"
)

type node struct {
	data int
	next *node
}

func insertAtBack(head *node, value int) *node {
	if head == nil {
		tempHead := new(node)
		tempHead.data = value
		zap.L().Debug("Insert at Root Node")
		return tempHead
	}
	var tempNode *node = head
	for tempNode.next != nil {
		tempNode = tempNode.next
	}
	newNode := new(node)
	newNode.data = value
	tempNode.next = newNode
	zap.L().Debug("Insert at End of Linked List")
	return head
}

func printList(head *node) {

	for head != nil {
		fmt.Println(head.data)
		head = head.next
		zap.L().Info("Printing Merged Linked List", zap.String("statusCode", "Success"))
	}
}
