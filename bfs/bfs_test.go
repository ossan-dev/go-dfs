package bfs_test

import (
	"testing"

	"github.com/ossan-dev/go-bfs/bfs"
	"github.com/stretchr/testify/require"
)

var g *bfs.Graph = &bfs.Graph{}

func BenchmarkWalkFromNode(b *testing.B) {
	b.ReportAllocs()
	expected := []string{"operating systems", "data structures", "computer organization", "discrete math", "intro to programming"}
	courses := make([]string, 0, numberOfCourses)
	queue := make([]string, 0, numberOfCourses)
	for b.Loop() {
		result := g.WalkFromNodeSlice("networks", courses, queue) // slices
		// result := g.WalkFromNodeMap("networks", courses, queue) // maps
		require.ElementsMatch(b, expected, result)
	}
}
