package tree

// Todo: NewNode creates new tree node
func NewNode(val int, left, right *Node) *Node {
	return &Node{
		Val:   val,
		left:  left,
		Right: right,
	}
}

// Todo: Node stores tree structure like data
type Node struct {
	Val   int
	left  *Node
	Right *Node
}

func IsBinarySearchTree(node *Node) bool {

	// if node.left != nil && node.left.Val >= node.Val {
	// 	return false
	// }

	// if node.Right != nil && node.Right.Val <= node.Val {
	// 	return false
	// }

	min := -9999
	max := 9999
	return isBinarySearchTree(node, min, max)
}

func isBinarySearchTree(root *Node, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isBinarySearchTree(root.left, min, root.Val) && isBinarySearchTree(root.Right, root.Val, max)
}
