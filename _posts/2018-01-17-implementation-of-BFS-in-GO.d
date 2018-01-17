---
published: true
author: surasnayak
layout: post
title: Implementation of Breadth First Search in GO
---


## Breadth First Search

- Breadth-first search (BFS) is an algorithm for traversing or searching tree or graph data structures. 
- It starts at the tree root (or some arbitrary node of a graph, sometimes referred to as a 'search key') and explores the neighbor nodes first, before moving to the next level neighbours.

### Traversing child nodes :

- A graph can contain cycles, which may bring you to the same node again while traversing the graph.
- To avoid processing of same node again, use a boolean array which marks the node after it is processed. 
- While visiting the nodes in the layer of a graph, store them in a manner such that you can traverse the corresponding child nodes in a similar order.

- To make this process easy, use a queue to store the node and mark it as 'visited' until all its neighbours (vertices that are directly connected to it) are marked.
- The queue follows the First In First Out (FIFO) queuing method, and therefore, the neigbors of the node will be visited in the order in which they were inserted in the node i.e. the node that was inserted first will be visited first, and so on.

### Applications :

1. To determine the level of each node in the given tree.
2. . 0-1 BFS

### Complexity

- The time complexity of BFS is O(V + E), where V is the number of nodes and E is the number of edges.

## Code for implementaion of Stack data structure in GO :

```go

package main

func bfs(start int, nodes map[int][]int, fn func (int)) {
    frontier := []int{start}
    visited := map[int]bool{}
    next := []int{}

    for 0 < len(frontier) {
        next = []int{}
        for _, node := range frontier {
            visited[node] = true
            fn(node)
            for _, n := range bfs_frontier(node, nodes, visited) {
                next = append(next, n)
            }
        }
        frontier = next
    }
}

func bfs_frontier(node int, nodes map[int][]int, visited map[int]bool) []int {
    next := []int{}
    iter := func (n int) bool { _, ok := visited[n]; return !ok }
    for _, n := range nodes[node] {
        if iter(n) {
            next = append(next, n)
        }
    }
    return next
}

```
