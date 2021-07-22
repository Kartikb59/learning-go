package btree

import (
	"fmt"

	"go.uber.org/zap"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func InsertNode(root *Node, value int) *Node {
	if root == nil {
		tempRoot := new(Node)
		tempRoot.data = value
		zap.L().Debug("Make a Root Node")
		return tempRoot
	}
	if root.data > value {
		zap.L().Debug("Traversing Left Side of the root node")
		root.left = InsertNode(root.left, value)
	} else {
		zap.L().Debug("Traversing Right Side of the root node")
		root.right = InsertNode(root.right, value)
	}
	return root
}

func Inorder(root *Node) {
	zap.L().Info("Printed Inorder Traversal of BST", zap.String("statusCode", "success"))
	if root == nil {
		return
	}
	Inorder(root.left)
	zap.L().Debug("Printing root node data according to inorder traversal")
	fmt.Println(root.data)
	Inorder(root.right)
}
