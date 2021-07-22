package main

import "go.uber.org/zap"

func mergeLL(head1 *node, head2 *node) *node {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}
	var mergedHead *node = new(node)
	if head1.data <= head2.data {
		mergedHead = head1
		head1 = head1.next
	} else {
		mergedHead = head2
		head2 = head2.next
	}

	var mergedTail *node = mergedHead
	for head1 != nil && head2 != nil {
		var temp *node = new(node)
		if head1.data <= head2.data {
			temp = head1
			head1 = head1.next
			zap.L().Debug("First LL Node value is smaller")
		} else {
			temp = head2
			head2 = head2.next
			zap.L().Debug("Second LL Node value is smaller")
		}
		mergedTail.next = temp
		mergedTail = mergedTail.next
	}

	if head1 != nil {
		mergedTail.next = head1
	} else if head2 != nil {
		mergedTail.next = head2
	}
	return mergedHead
}
