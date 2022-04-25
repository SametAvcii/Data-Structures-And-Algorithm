package main

type AvlTree struct {
	height, value int
	right, left   *AvlTree
}

func (n *AvlTree) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}
func (n *AvlTree) BalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.height - n.right.height
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func (n *AvlTree) UpdateHeight() {
	n.height = max(n.left.Height(), n.right.height) + 1
}
func (n *AvlTree) RightRotate() *AvlTree {
	y := n.left
	z := y.left

	y.right = n
	n.left = z
	n.UpdateHeight()
	y.UpdateHeight()

	return y
}
func (n *AvlTree) LeftRotate() *AvlTree {
	y := n.right
	t := y.left
	y.left = n
	n.right = t

	n.UpdateHeight()
	y.UpdateHeight()

	return y
}
