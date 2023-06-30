package main

import (
    "testing"
    "strings"
)

func TestDFS(t *testing.T) {
    // Create a graph
    graph := Graph{}

    // Create vertices
    v1 := &Vertex{value: 1}
    v2 := &Vertex{value: 2}
    v3 := &Vertex{value: 3}
    v4 := &Vertex{value: 4}
    v5 := &Vertex{value: 5}

    // Add vertices to the graph
    graph.addVertex(v1)
    graph.addVertex(v2)
    graph.addVertex(v3)
    graph.addVertex(v4)
    graph.addVertex(v5)

    // Add edges between vertices
    v1.addAdjacent(v2)
    v1.addAdjacent(v3)
    v2.addAdjacent(v4)
    v3.addAdjacent(v5)

    // Perform depth first search
    sb := strings.Builder{}
    dfs(v1, &sb)
    output := sb.String()
    expected := "1 2 4 3 5 "

    if output != expected {
        t.Errorf("Expected %s, but got %s", expected, output)
    }
}
