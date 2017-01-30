package RBTree

type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
	isRed       bool
}

// Get returns TreeNode with "search" value
// return nil is "search" not found
func (t *Tree) Get(search int) *TreeNode {
	if t.root == nil {
		return nil
	}
	current := t.root
	for current != nil {
		if current.Value == search {
			return current
		}
		if current.Value > search {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return nil
}

// rotateRight perform right rotation
func (t *Tree) rotateRight(node *TreeNode) *TreeNode {
	x := node.Left
	node.Left = x.Right
	x.Right = node
	x.isRed = x.Right.isRed
	x.Right.isRed = true
	return x
}

// rotateLeft perform left rotation
func (t *Tree) rotateLeft(node *TreeNode) *TreeNode {
	x := node.Right
	node.Right = x.Left
	x.Left = node
	x.isRed = x.Left.isRed
	x.Left.isRed = true
	return x
}

// flipColor flip red black colors
func (t *Tree) flipColor(node *TreeNode) {
	node.isRed = !node.isRed
	node.Left.isRed = !node.Left.isRed
	node.Right.isRed = !node.Right.isRed
}

// put new value to Tree
func (t *Tree) put(root *TreeNode, value int) *TreeNode {
	if root == nil {
		result := &TreeNode{Value: value}
		result.isRed = true
		return result
	}
	if value < root.Value {
		root.Left = t.put(root.Left, value)
	}
	if value > root.Value {
		root.Right = t.put(root.Right, value)
	}

	// Fix tree structure
	if root.Right != nil && root.Right.isRed && root.Left != nil && !root.Left.isRed {
		root = t.rotateLeft(root)
	}
	if root.Left != nil && root.Left.isRed && root.Left.Left != nil && root.Left.Left.isRed {
		root = t.rotateRight(root)
	}
	if root.Left != nil && root.Left.isRed && root.Right != nil && root.Right.isRed {
		t.flipColor(root)
	}
	return root
}

// Put new value to Tree
func (t *Tree) Put(value int) {
	t.root = t.put(t.root, value)
	t.root.isRed = false
}

// getOrdered returns ordered values in Tree
func (t *Tree) getOrdered(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := t.getOrdered(root.Left)
	result = append(result, root.Value)
	result = append(result, t.getOrdered(root.Right)...)
	return result
}

// GetOrdered returns ordered values in Tree
func (t *Tree) GetOrdered() []int {
	return t.getOrdered(t.root)
}
