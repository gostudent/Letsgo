---
published: true
author: surasnayak
layout: post
title: Implementation of Trie Datastructure in GO
---

## Trie Datastructure

- Tries are an extremely special and useful data-structure that are based on the prefix of a string. 
- They are used to represent the “Retrieval” of data and thus the name Trie.

### Prefix : What is prefix ?

The prefix of a string is nothing but any n letters (n≤|S|)  that can be considered beginning strictly from the starting of a string. 

-For example , the word “abacaba” has the following prefixes:

a
ab
aba
abac
abaca
abacab

- A Trie is a special data structure used to store strings that can be visualized like a graph. 
- It consists of nodes and edges. 
- Each node consists of at max 26 children and edges connect each parent node to its children. 
- These 26 pointers are nothing but pointers for each of the 26 letters of the English alphabet A separate edge is maintained for every edge.

- Strings are stored in a top to bottom manner on the basis of their prefix in a trie. 
- All prefixes of length 1 are stored at until level 1, all prefixes of length 2 are sorted at until level 2 and so on.


### Inserting into a Trie :

- Inserting a key into Trie is simple approach. 
- Every character of input key is inserted as an individual Trie node. 
- Note that the children is an array of pointers (or references) to next level trie nodes. 
- The key character acts as an index into the array children. 
- If the input key is new or an extension of existing key, we need to construct non-existing nodes of the key, and mark end of word for last node. 
- If the input key is prefix of existing key in Trie, we simply mark the last node of key as end of word. The key length determines Trie depth.


### Searching for a key in Trie :


- Searching for a key is similar to insert operation, however we only compare the characters and move down. 
- The search can terminate due to end of string or lack of key in trie. 
- In the former case, if the isEndofWord field of last node is true, then the key exists in trie. 
- In the second case, the search terminates without examining all the characters of key, since the key is not present in trie.


## Code for implementaion of Trie data structure in GO :

```go

// Package trie provides Trie data structures in golang.
//
// Wikipedia: https://en.wikipedia.org/wiki/Trie
package trie

// Node represents each node in Trie.
type Node struct {
	children map[rune]*Node // map children nodes
	isLeaf   bool           // current node value
}

// NewNode creates a new Trie node with initialized
// children map.
func NewNode() *Node {
	n := &Node{}
	n.children = make(map[rune]*Node)
	n.isLeaf = false
	return n
}

// Insert inserts words at a Trie node.
func (n *Node) Insert(s string) {
	curr := n
	for _, c := range s {
		next, ok := curr.children[c]
		if !ok {
			next = NewNode()
			curr.children[c] = next
		}
		curr = next
	}
	curr.isLeaf = true
}

// Find finds words at a Trie node.
func (n *Node) Find(s string) bool {
	curr := n
	for _, c := range s {
		next, ok := curr.children[c]
		if !ok {
			return false
		}
		curr = next
	}
	return true
}

```