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
	return false
}
