package bfs

import (
	"fmt"
	"slices"
)

type Graph struct {
	vertices []*vertex
}

type vertex struct {
	key      string
	adjacent []*vertex
}

func (g *Graph) AddVertex(k string) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &vertex{key: k})
}

func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
		return
	}
	if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
		return
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

func (g *Graph) WalkFromNodeSlice(startNode string, courses, queue []string) []string {
	queue = append(queue, startNode)
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if currentNode != startNode {
			courses = append(courses, currentNode)
		}
		// adjacents to the current node
		for _, adjacent := range g.getVertex(currentNode).adjacent {
			if !slices.Contains(courses, adjacent.key) {
				queue = append(queue, adjacent.key)
			}
		}
	}
	return courses
}

func (g *Graph) WalkFromNodeMap(startNode string, courses, queue []string) []string {
	visitedNodes := make(map[string]struct{})
	queue = append(queue, startNode)
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		visitedNodes[currentNode] = struct{}{}
		// adjacents to the current node
		for _, adjacent := range g.getVertex(currentNode).adjacent {
			if _, isFound := visitedNodes[adjacent.key]; !isFound {
				queue = append(queue, adjacent.key)
			}
		}
	}

	for f := range visitedNodes {
		if f == startNode {
			continue
		}
		courses = append(courses, f)
	}
	return courses
}

func (g *Graph) getVertex(k string) *vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*vertex, k string) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
