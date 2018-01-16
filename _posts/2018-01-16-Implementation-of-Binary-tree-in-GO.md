---
published: true
author: surasnayak
layout: post
title: Implementation of Binary Tree Datastructure in GO
---

## Binary Tree Datastructure

- We extend the concept of linked data structures to structure containing nodes with more than one self-referenced field. 
- A binary tree is made of nodes, where each node contains a "left" reference, a "right" reference, and a data element. The topmost node in the tree is called the root.
- Every node (excluding a root) in a tree is connected by a directed edge from exactly one other node. This node is called a parent. 
- On the other hand, each node can be connected to arbitrary number of nodes, called children. 
- Nodes with no children are called leaves, or external nodes. 
- Nodes which are not leaves are called internal nodes. 
- Nodes with the same parent are called siblings.

## Important terminologies :

- The depth of a node is the number of edges from the root to the node.
- The height of a node is the number of edges from the node to the deepest leaf.
- The height of a tree is a height of the root.
- A full binary tree.is a binary tree in which each node has exactly zero or two children.
- A complete binary tree is a binary tree, which is completely filled, with the possible exception of the bottom level, which is filled from left to right.

## Binary Tree Representation :

- A tree is represented by a pointer to the topmost node in tree. If the tree is empty, then value of root is NULL.

- A Tree node contains following parts.

1. Data
2. Pointer to left child
3. Pointer to right child

## Tree Traversals :

### Algorithm for Inorder Tree Traversal :

- Traverse the left subtree, i.e., call Inorder(left-subtree)
- Visit the root.
- Traverse the right subtree, i.e., call Inorder(right-subtree)

### Algorithm for Preorder Tree Traversal :

- Visit the root.
- Traverse the left subtree, i.e., call Preorder(left-subtree)
- Traverse the right subtree, i.e., call Preorder(right-subtree) 

### Algorithm for Postorder Tree Traversal :

- Traverse the left subtree, i.e., call Postorder(left-subtree)
- Traverse the right subtree, i.e., call Postorder(right-subtree)
- Visit the root.

## Code for implementaion of Binary Tree data structure in GO :

```go

package binarytree

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

type btree struct {
	root *node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(val int) *node {
	n := &node{val, nil, nil}
	return n
}

func inorder(n *node) {
	if n == nil {
		return
	}
	inorder(n.left)
	fmt.Print(n.val, " ")
	inorder(n.right)
}

func preorder(n *node) {
	if n == nil {
		return
	}
	fmt.Print(n.val, " ")
	inorder(n.left)
	inorder(n.right)
}

func postorder(n *node) {
	if n == nil {
		return
	}
	inorder(n.left)
	inorder(n.right)
	fmt.Print(n.val, " ")
}

func levelorder(root *node) {
	var q []*node // queue
	var n *node   // temporary node

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		fmt.Print(n.val, " ")
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

// helper function for t.depth
func _calculate_depth(n *node, depth int) int {
	if n == nil {
		return depth
	}
	return max(_calculate_depth(n.left, depth+1), _calculate_depth(n.right, depth+1))
}

func (t *btree) depth() int {
	return _calculate_depth(t.root, 0)
}

```