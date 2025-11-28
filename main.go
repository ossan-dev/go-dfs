package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type graph struct {
	vertices []*vertex
}

type vertex struct {
	key      string
	adjacent []*vertex
}

func (g *graph) addVertex(k string) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &vertex{key: k})
}

func (g *graph) addEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
		return
	}
	if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
		return
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

func (g *graph) walkFromNode(startNode string) []string {
	var visitedNodes []string
	queue := []string{startNode}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		visitedNodes = append(visitedNodes, currentNode)
		// adjacents to the current node
		for _, adjacent := range g.getVertex(currentNode).adjacent {
			flag := 1
			for _, v := range visitedNodes {
				if adjacent.key == v {
					flag = 0
					break
				}
			}
			if flag == 1 {
				queue = append(queue, adjacent.key)
			}
		}
	}
	return visitedNodes[1:]
}

func (g *graph) getVertex(k string) *vertex {
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

func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %q :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %q ", v.key)
		}
	}
	fmt.Println()
}

func main() {
	file, err := os.Open("course.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var courses map[string][]string
	if err := json.NewDecoder(file).Decode(&courses); err != nil {
		panic(err)
	}
	// for k, v := range courses {
	// 	fmt.Printf("%q - %q\n", k, v)
	// }

	// graph
	// algorithmsSubCourses, isFound := courses["algorithms"]
	// if !isFound {
	// 	return
	// }
	g := &graph{}
	for k := range courses {
		g.addVertex(k)
	}
	for k, v := range courses {
		for _, vv := range v {
			g.addEdge(k, vv)
		}
	}
	// g.print()

	// leaf node
	// fmt.Printf("linear algebra: %q\n", g.walkFromNode("linear algebra"))
	// root node
	fmt.Printf("algorithms: %q\n", g.walkFromNode("algorithms"))

}
