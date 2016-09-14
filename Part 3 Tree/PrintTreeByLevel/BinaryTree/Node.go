package BinaryTree

import "math"

// Node in Binary Tree
// Data      interface{}
// Parent    *Node
// LeftNode  *Node
// RightNode *Node
// Size      int64
// Height    int64
type Node struct {
	Data      interface{}
	Parent    *Node
	LeftNode  *Node
	RightNode *Node
	Size      int64
	Height    int64
}

/*
                      1
                   /     \
                  2       3
                 / \      /
                4   3    6
               / \   \  / \
              7   8  9  10 11
             /     \
            12      13
                   /
                  14
Output:
1
2 3
4 3 6
7 8 9 10 11
12 13
14
*/

// NewNode return a Bin Tree
// @param  interface
// @return *Node
func NewNode(root interface{}) *Node {
	return &Node{Data: root, Size: 1}
}

// GetData return data on Node
// @return interface{}
func (node *Node) GetData() interface{} {
	if node == nil {
		return nil
	}
	return node.Data
}

// SetData set data on Node
// @param interface{}
func (node *Node) SetData(value interface{}) {
	node.Data = value
}

// HasParent return Has Parent Node
// @return bool
func (node *Node) HasParent() bool {
	return node.Parent != nil
}

// GetParent return Parent Node on Node
// @return *Node
func (node *Node) GetParent() *Node {
	if !node.HasParent() {
		return nil
	}
	return node.Parent
}

// SetParent set Parent Node on Node.
// @param *Node
func (node *Node) SetParent(pn *Node) {
	node.Parent = pn
}

// CutOffParent cut off Parent on Node.
func (node *Node) CutOffParent() {
	if !node.HasParent() {
		return
	}

	if node.IsLeftChild() {
		node.Parent.LeftNode = nil
	} else {
		node.Parent.RightNode = nil
	}

	node.Parent = nil
	node.Parent.SetHeight()
	node.Parent.SetSize()
}

// HasLeftChild return Has Left Node
// @return bool
func (node *Node) HasLeftChild() bool {
	return node.LeftNode != nil
}

// GetLeftNode return Left Node on Node
// @return *Node
func (node *Node) GetLeftNode() *Node {
	if !node.HasLeftChild() {
		return nil
	}
	return node.LeftNode
}

// SetLeftNode set Left Node on Node and Return old  Left Node.
// @param ln *Node
// @return oldLN *Node
func (node *Node) SetLeftNode(ln *Node) *Node {
	oldLN := node.LeftNode
	if node.HasLeftChild() {
		node.LeftNode.CutOffParent()
	}
	if ln != nil {
		ln.CutOffParent()
		node.LeftNode = ln
		ln.SetParent(node)
		node.SetHeight()
		node.SetSize()
	}
	return oldLN
}

// HasRightChild return Has Right Node
// @return bool
func (node *Node) HasRightChild() bool {
	return node.RightNode != nil
}

// GetRightNode return Right Node on Node
// @return *Node
func (node *Node) GetRightNode() *Node {
	if !node.HasRightChild() {
		return nil
	}
	return node.RightNode
}

// SetRightNode set Right Node on Node and Return old  Right Node.
// @param rn *Node
// @return oldLN *Node
func (node *Node) SetRightNode(rn *Node) *Node {
	oldRN := node.RightNode
	if node.HasRightChild() {
		node.RightNode.CutOffParent()
	}
	if rn != nil {
		rn.CutOffParent()
		node.RightNode = rn
		rn.SetParent(node)
		node.SetHeight()
		node.SetSize()
	}
	return oldRN
}

// IsLeaf return is Leaf Node.
// @return bool
func (node *Node) IsLeaf() bool {
	return !node.HasLeftChild() && !node.HasRightChild()
}

// IsLeftChild return is Left Child.
// @return bool
func (node *Node) IsLeftChild() bool {
	return node.HasParent() && node == node.Parent.LeftNode
}

// IsRightChild return is Right Child.
// @return bool
func (node *Node) IsRightChild() bool {
	return node.HasParent() && node == node.Parent.RightNode
}

// GetHeight return Height on Node.
// @return int64
func (node *Node) GetHeight() int64 {
	return node.Height
}

// SetHeight set Height on Node.
func (node *Node) SetHeight() {
	var newH int64
	if node.HasLeftChild() {
		newH = int64(math.Max(float64(newH), float64(node.LeftNode.GetHeight()+1)))
	}
	if node.HasRightChild() {
		newH = int64(math.Max(float64(newH), float64(node.RightNode.GetHeight()+1)))
	}

	if newH == node.Height {
		return
	}
	node.Height = newH

	if node.HasParent() {
		node.GetParent().SetHeight()
	}
}

// GetSize return Size on Node.
// @return int64
func (node *Node) GetSize() int64 {
	return node.Size
}

// SetSize set Size on Node.
func (node *Node) SetSize() {
	node.Size = 1
	if node.HasLeftChild() {
		node.Size += node.GetLeftNode().GetSize()
	}
	if node.HasRightChild() {
		node.Size += node.GetRightNode().GetSize()
	}

	if node.HasParent() {
		node.Parent.SetSize()
	}
}
