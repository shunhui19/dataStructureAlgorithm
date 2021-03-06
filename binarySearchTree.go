package dataStructureAlgorithm

import (
	"fmt"
)

// BinarySearchTree 二叉查找树, 主要应用于查找方面.
// 特点：对于任意一个节点X，它的左子树中的 "所有的值" 小于X的值，它的右子树中的 "所有的值" 大于X的值,
// 这里是指所有节点的值
type BinarySearchTree struct {
	// Value 树节点存储的值
	Value int
	// left 左子树
	left *BinarySearchTree
	// right 右子树
	right *BinarySearchTree
}

// MakeEmpty 清空树.
func (b *BinarySearchTree) MakeEmpty() {
	b.Value, b.left, b.right = 0, nil, nil
}

// Find 返回元素值的树.
func (b *BinarySearchTree) Find(data int) *BinarySearchTree {
	// 如果相等，返回
	if b.Value == data {
		return b
	}

	// 如果大于根节点值, 递归遍历右子树
	if data > b.Value {
		for b.right != nil {
			return b.right.Find(data)
		}

		// 非递归遍历
		//right := (*b).right
		//for right != nil {
		//	if right.Value == data {
		//		return right
		//	}
		//	right = right.right
		//}
	}

	// 如果小于根节点值, 递归遍历左子树
	if data < b.Value {
		for b.left != nil {
			return b.left.Find(data)
		}
	}

	// 不存在data值的节点
	return nil
}

// FindMin 获取最小的值.
// 采用非递归方式.
func (b *BinarySearchTree) FindMin() int {
	if b.IsEmpty() {
		return 0
	}

	// 值拷贝传递
	left := (*b).left
	for left.left != nil {
		left = left.left
	}

	return left.Value
}

// FindMax 获取最大的值.
// 采用递归方式.
func (b *BinarySearchTree) FindMax() int {
	if b.IsEmpty() {
		return 0
	}

	for b.right != nil {
		return b.right.FindMax()
	}

	return b.Value
}

// IsEmpty 是否为空树.
func (b *BinarySearchTree) IsEmpty() bool {
	return b.Value == 0 && b.left == nil && b.right == nil
}

// Insert 插入.
// 空树：只有根节点，并且值为0， 没有左子树也没有右子树.
func (b *BinarySearchTree) Insert(data int) *BinarySearchTree {
	if b.Value == 0 && b.left == nil && b.right == nil {
		b.Value = data
		return nil
	}

	// 值已存在直接返回
	if b.Value == data {
		return nil
		// 值小于当前节点的值
	} else if data < b.Value {
		if b.left == nil {
			b.left = &BinarySearchTree{
				Value: data,
				left:  nil,
				right: nil,
			}
			return nil
		} else {
			b.left.Insert(data)
		}
		// 值大于当前节点的值
	} else {
		if data > b.Value {
			if b.right == nil {
				b.right = &BinarySearchTree{
					Value: data,
					left:  nil,
					right: nil,
				}
				return nil
			} else {
				b.right.Insert(data)
			}
		}
	}
	return nil
}

// PreOrderTraversalRecursion 递归先序遍历.
// 先遍历根节点，然后遍历左子树，再遍历右子树.
func (b *BinarySearchTree) PreOrderTraversalRecursion() {
	if b.IsEmpty() {
		return
	}

	fmt.Printf("%v ", b.Value)
	if b.left != nil {
		b.left.PreOrderTraversalRecursion()
	}
	if b.right != nil {
		b.right.PreOrderTraversalRecursion()
	}
}

// MidOrderTraversalRecursion 递归中序遍历.
// 先遍历左子树，然后遍历根节点，再遍历右子树.
func (b *BinarySearchTree) MidOrderTraversalRecursion() {
	if b.IsEmpty() {
		return
	}

	if b.left != nil {
		b.left.MidOrderTraversalRecursion()
	}
	fmt.Printf("%v ", b.Value)
	if b.right != nil {
		b.right.MidOrderTraversalRecursion()
	}
}

// PostOrderTraversalRecursion 递归后序遍历.
// 先遍历左子树，然后遍历右子树，再遍历根节点.
func (b *BinarySearchTree) PostOrderTraversalRecursion() {
	if b.IsEmpty() {
		return
	}

	if b.left != nil {
		b.left.PostOrderTraversalRecursion()
	}
	if b.right != nil {
		b.right.PostOrderTraversalRecursion()
	}
	fmt.Printf("%v ", b.Value)
}

// Delete 删除节点, 根据要删除节点子节点的个数分三种情况：
// 1.要删除的节点是叶子节点，直接设置删除节点的父节点指向删除节点的指针为nil
// 2.要删除的节点只有一个节点，设置删除节点的父节点指向删除节点的指针指向"要删除节点的子节点"
// 3.要删除的节点有两个节点，从删除节点的右子树查询最小节点值替换到要删除节点的值，同时删除最小节点
func (b *BinarySearchTree) Delete(data int) {
	// 删除节点的父节点
	pp := b
	// 待删除的节点
	p := b

	for p.Value != 0 && p.Value != data {
		pp = p
		if data > p.Value {
			p = p.right
		} else {
			p = p.left
		}
	}
	// 没有找到要删除的节点
	if p == nil {
		return
	}

	// 第一种情况, 判断要删除节点是位于左节点还是右节点
	if p.left == nil && p.right == nil {
		if pp.left == p {
			pp.left = nil
		} else {
			pp.right = nil
		}
	}

	// 第二种情况，只有一个节点
	if (p.left == nil && p.right != nil) || (p.left != nil && p.right == nil) {
		if pp.left == p {
			pp.left = p.left
		} else {
			pp.right = p.right
		}
	}

	// 第三种情况, 删除节点有两个子节点
	if p.left != nil && p.right != nil {
		// minP的父节点
		minPP := p
		// 查找子右节点的最小节点
		minP := p.right
		for minP.left != nil {
			minPP = minP
			minP = minP.left
		}
		// 把右子树中最小节点值替换到要删除节点的值
		p.Value = minP.Value

		// 删除minP节点
		// 如果最小节点是左叶子节点, 直接删除最小节点
		if minPP.left == minP {
			minPP.left = nil
			// 最小节点不是左叶子节点, 说明当前minP只包含右子树
		} else {
			minPP.right = minP.right
		}
	}
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Value: 0,
		left:  nil,
		right: nil,
	}
}
