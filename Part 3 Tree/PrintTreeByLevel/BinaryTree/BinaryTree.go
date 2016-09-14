package BinaryTree

import "container/list"

// BinaryTree ...
type BinaryTree struct {
	Root   *Node
	Height int64
	Size   int64
}

// NewBinaryTree return a BinaryTree with root Node.
func NewBinaryTree(root *Node) *BinaryTree {
	return BinaryTree{Root: root}
}

// GetSize return root Size.
// @return int64
func (bt *BinaryTree) GetSize() int64 {
	return bt.Root.Size
}

// GetHeight return root Height.
// @return int64
func (bt *BinaryTree) GetHeight() int64 {
	return bt.Root.Height
}

// IsEmpty return BinaryTree is Empty.
// @return bool
func (bt *BinaryTree) IsEmpty() bool {
	return bt.Root != nil
}

// GetRoot return Root on BinaryTree
// @return *Node
func (bt *BinaryTree) GetRoot() *Node {
	return bt.Root
}

// Find return first Node which Data Equal value.
// @param value interface{}
// @return *Node
func (bt *BinaryTree) Find(value interface{}) *Node {
	if bt.Root == nil {
		return nil
	}
	p := bt.Root
	return isEqual(value, p)
}

func isEqual(value interface{}, node *Node) *Node {
	if value == node.GetData() {
		return node
	}

	if node.HasLeftChild() {
		ln := isEqual(value, node.GetLeftNode())
		if ln != nil {
			return ln
		}
	}

	if node.HasRightChild() {
		rn := isEqual(value, node.GetRightNode())
		if rn != nil {
			return rn
		}
	}
	return nil
}

// ProOrder return a List By ProOrder
// @return *list.List
func (bt *BinaryTree) ProOrder() *list.List {
	traversal := list.New()
	proOrder(bt.Root, traversal)
	return traversal
}

func proOrder(n *Node, l *list.List) {
	if n == nil {
		return
	}
	l.PushBack(n)
	proOrder(n.GetLeftNode(), l)
	proOrder(n.GetRightNode(), l)
}

// InOrder return a List By InOrder
// @return *list.List
func (bt *BinaryTree) InOrder() *list.List {
	traversal := list.New()
	inOrder(bt.Root, traversal)
	return traversal
}

func inOrder(n *Node, l *list.List) {
	if n == nil {
		return
	}
	inOrder(n.GetLeftNode(), l)
	l.PushBack(n)
	inOrder(n.GetRightNode(), l)
}

// PostOrder return a List By PostOrder
// @return *list.List
func (bt *BinaryTree) PostOrder() *list.List {
	traversal := list.New()
	postOrder(bt.Root, traversal)
	return traversal
}

func postOrder(n *Node, l *list.List) {
	if n == nil {
		return
	}
	postOrder(n.GetLeftNode(), l)
	postOrder(n.GetRightNode(), l)
	l.PushBack(n)
}
