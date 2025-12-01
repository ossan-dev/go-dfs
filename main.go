package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ossan-dev/go-bfs/bfs"
)

// adjacent lists okay

func main() {
	// load data from JSON file
	file, err := os.Open("course.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var courses map[string][]string
	if err := json.NewDecoder(file).Decode(&courses); err != nil {
		panic(err)
	}
	// build the graph from the map
	g := &bfs.Graph{}
	for k := range courses {
		g.AddVertex(k)
	}
	for k, v := range courses {
		for _, vv := range v {
			g.AddEdge(k, vv)
		}
	}

	// test the graph
	coursesToAttend := make([]string, 0, len(courses))
	queue := make([]string, 0, len(courses))
	toPrint := g.WalkFromNodeSlice("networks", coursesToAttend, queue)
	fmt.Printf("slice\tnetworks: %q\n", toPrint)
	coursesToAttend2 := make([]string, 0, len(courses))
	toPrint = g.WalkFromNodeMap("networks", coursesToAttend2, queue)
	fmt.Printf("map\tnetworks: %q\n", toPrint)
}
