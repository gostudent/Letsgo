---
published: true
author: surasnayak
layout: post
title: Implementation of Topological Sort in GO
---

## Topological Sort :

- Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge uv, vertex u comes before v in the ordering. 
- Topological Sorting for a graph is not possible if the graph is not a DAG.


### Topological Sorting vs Depth First Traversal (DFS):

- In DFS, we print a vertex and then recursively call DFS for its adjacent vertices. 
- In topological sorting, we need to print a vertex before its adjacent vertices.

### Algorithm of Bucket Sort :

 - In topological sorting, we use a temporary stack. We don’t print the vertex immediately, we first recursively call topological sorting for all its adjacent vertices, then push it to a stack. 
 - Finally, print contents of stack. Note that a vertex is pushed to stack only when all of its adjacent vertices (and their adjacent vertices and so on) are already in stack.

 ## Code for implementaion of Topological Sort in GO :

 ```go

package topsort

import (
	"fmt"
	"strings"
)

type Graph struct {
	nodes map[string]node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]node),
	}
}

func (g *Graph) AddNode(name string) {
	if !g.ContainsNode(name) {
		g.nodes[name] = make(node)
	}
}

func (g *Graph) AddEdge(from string, to string) error {
	f, ok := g.nodes[from]
	if !ok {
		return fmt.Errorf("Node %q not found", from)
	}
	_, ok = g.nodes[to]
	if !ok {
		return fmt.Errorf("Node %q not found", to)
	}

	f.addEdge(to)
	return nil
}

func (g *Graph) ContainsNode(name string) bool {
	_, ok := g.nodes[name]
	return ok
}

func (g *Graph) TopSort(name string) ([]string, error) {
	results := newOrderedSet()
	err := g.visit(name, results, nil)
	if err != nil {
		return nil, err
	}
	return results.items, nil
}

func (g *Graph) visit(name string, results *orderedset, visited *orderedset) error {
	if visited == nil {
		visited = newOrderedSet()
	}

	added := visited.add(name)
	if !added {
		index := visited.index(name)
		cycle := append(visited.items[index:], name)
		return fmt.Errorf("Cycle error: %s", strings.Join(cycle, " -> "))
	}

	n := g.nodes[name]
	for _, edge := range n.edges() {
		err := g.visit(edge, results, visited.copy())
		if err != nil {
			return err
		}
	}

	results.add(name)
	return nil
}

type node map[string]bool

func (n node) addEdge(name string) {
	n[name] = true
}

func (n node) edges() []string {
	var keys []string
	for k := range n {
		keys = append(keys, k)
	}
	return keys
}

type orderedset struct {
	indexes map[string]int
	items   []string
	length  int
}

func newOrderedSet() *orderedset {
	return &orderedset{
		indexes: make(map[string]int),
		length:  0,
	}
}

func (s *orderedset) add(item string) bool {
	_, ok := s.indexes[item]
	if !ok {
		s.indexes[item] = s.length
		s.items = append(s.items, item)
    s.length++
	}
	return !ok
}

func (s *orderedset) copy() *orderedset {
	clone := newOrderedSet()
	for _, item := range s.items {
		clone.add(item)
	}
	return clone
}

func (s *orderedset) index(item string) int {
	index, ok := s.indexes[item]
	if ok {
		return index
	}
	return -1
}

 ```
