package datastructure

func LCA(root *AvlNode, n1, n2 int) *AvlNode {

	if root == nil {
		return nil
	}

	if n1 < root.key && n2 < root.key {
		return LCA(root.left, n1, n2)
	}

	if n1 > root.key && n2 > root.key {
		return LCA(root.right, n1, n2)
	}

	return root
}

func LCAIterative(root *AvlNode, n1, n2 int) *AvlNode {

	for root != nil {
		if n1 < root.key && n2 < root.key {
			root = root.left
		} else if n1 > root.key && n2 > root.key {
			root = root.right
		} else {
			break
		}
	}

	return root
}
