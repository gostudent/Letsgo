---
published: true
author: surasnayak
layout: post
title: Implementation of Depth First Search in GO
---


## Depth First Search :

- The DFS algorithm is a recursive algorithm that uses the idea of backtracking. 
- It involves exhaustive searches of all the nodes by going ahead, if possible, else by backtracking.

- Here, the word backtrack means that when you are moving forward and there are no more nodes along the current path, you move backwards on the same path to find nodes to traverse. 
- All the nodes will be visited on the current path till all the unvisited nodes have been traversed after which the next path will be selected.
- This recursive nature of DFS can be implemented using stacks.

### basic idea for implementation : 

- Pick a starting node and push all its adjacent nodes into a stack.
- Pop a node from stack to select the next node to visit and push all its adjacent nodes into a stack.
- Repeat this process until the stack is empty. However, ensure that the nodes that are visited are marked. 
- This will prevent you from visiting the same node more than once.
- If you do not mark the nodes that are visited and you visit the same node more than once, you may end up in an infinite loop.

### Handling disconnected graphs :

- Sometimes all the vertices may not be reachable from a given vertex. 
- To do complete DFS traversal of such graphs, we must call DFSUtil() for every vertex. 
- Also, before calling DFSUtil(), we should check if it is already printed by some other call of DFSUtil().

### Applications

- To find connected components using DFS.

### Time complexity :

- Time complexity O(V+E).

## Code for implementaion of Depth First Search in GO :

```go

package main

func dfs(node int, nodes map[int][]int, fn func (int)) {
    dfs_recur(node, map[int]bool{}, fn)
}

func dfs_recur(node int, v map[int]bool, fn func (int)) {
    v[node] = true
    fn(node)
    for _, n := range nodes[node] {
        if _, ok := v[n]; !ok {
            dfs_recur(n, v, fn)
        }
    }
}

```

