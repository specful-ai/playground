package main

import (
    "fmt"
    "strings"
)

type Graph struct {
    vertices []*Vertex
}

type Vertex struct {
    value    int
    visited  bool
    adjacent []*Vertex
}

func (g *Graph) addVertex(v *Vertex) {
    g.vertices = append(g.vertices, v)
}

func (v *Vertex) addAdjacent(adj *Vertex) {
    v.adjacent = append(v.adjacent, adj)
}

func dfs(v *Vertex, sb *strings.Builder) {
    fmt.Fprintf(sb, "%d ", v.value)
    v.visited = true

    for _, adj := range v.adjacent {
        if !adj.visited {
            dfs(adj, sb)
        }
    }
}

func main() {
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
    fmt.Println("Depth First Search:")
    dfs(v1, &sb)
    fmt.Println(sb.String())
}
