package datastructure

import (
	"fmt"
)

type AvlNode struct {
	key    int
	height int
	left   *AvlNode
	right  *AvlNode
}

func (a *AvlNode) GetKey() int {
	return a.key
}

func GetBalance(root *AvlNode) int {

	if root == nil {
		return 0
	}

	return GetHeight(root.left) - GetHeight(root.right)
}

func GetHeight(root *AvlNode) int {
	if root == nil {
		return -1
	}

	return root.height
}

func leftRotate(root *AvlNode) *AvlNode {

	rightChild := root.right
	temp := rightChild.left

	rightChild.left = root
	root.right = temp

	root.height = 1 + max(GetHeight(root.left), GetHeight(root.right))
	rightChild.height = 1 + max(GetHeight(rightChild.left), GetHeight(rightChild.right))

	return rightChild
}

func rightRotate(root *AvlNode) *AvlNode {

	leftChild := root.left
	temp := leftChild.right

	leftChild.right = root
	root.left = temp

	root.height = 1 + max(GetHeight(root.left), GetHeight(root.right))
	leftChild.height = 1 + max(GetHeight(leftChild.left), GetHeight(leftChild.right))

	return leftChild
}

func Insert(key int, root *AvlNode) *AvlNode {

	if root == nil {
		return &AvlNode{
			key: key,
		}
	}

	if key < root.key {
		root.left = Insert(key, root.left)
	} else if key > root.key {
		root.right = Insert(key, root.right)
	} else {
		return root
	}

	root.height = 1 + max(GetHeight(root.left), GetHeight(root.right))
	balance := GetBalance(root)
	if balance < -1 && GetBalance(root.left) < 0 {
		return leftRotate(root)
	} else if balance < -1 && GetBalance(root.left) > 0 {
		root.left = rightRotate(root.left)
		return leftRotate(root)
	} else if balance > 1 && GetBalance(root.right) > 0 {
		return rightRotate(root)
	} else if balance > 1 && GetBalance(root.right) < 0 {
		root.right = leftRotate(root.right)
		return rightRotate(root)
	}

	return root
}

func Delete(key int, root *AvlNode) *AvlNode {

	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = Delete(key, root.left)
	} else if key > root.key {
		root.right = Delete(key, root.right)
	} else {

		if root.left == nil || root.right == nil {
			var temp *AvlNode
			if root.left == nil {
				temp = root.right
			} else {
				temp = root.left
			}

			if temp == nil {
				root = nil
			} else {
				*root = *temp
			}
		} else {

			nextBiggest := root.right
			for nextBiggest.left != nil {
				nextBiggest = nextBiggest.left
			}

			root.key = nextBiggest.key
			root.right = Delete(nextBiggest.key, root.right)

		}
	}

	if root == nil {
		return root
	}

	root.height = 1 + max(GetHeight(root.left), GetHeight(root.right))
	balance := GetBalance(root)
	if balance < -1 && GetBalance(root.left) < 0 {
		return leftRotate(root)
	} else if balance < -1 && GetBalance(root.left) > 0 {
		root.left = rightRotate(root.left)
		return leftRotate(root)
	} else if balance > 1 && GetBalance(root.right) > 0 {
		return rightRotate(root)
	} else if balance > 1 && GetBalance(root.right) < 0 {
		root.right = leftRotate(root.right)
		return rightRotate(root)
	}

	return root
}

func PreOrder(root *AvlNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.key)
	PreOrder(root.left)
	PreOrder(root.right)
}

// Time Complexity:
// Insert: O(log(n))
// Delete: O(log(n))
// PreOrder: O(n)
