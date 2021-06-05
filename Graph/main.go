package main

import "fmt"

func main() {
	g := &Graph{}
	for i := 0; i < 5; i++ {
		g.AddVertex(i)

	}
	g.AddEdge(2, 4)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	//fmt.Println(g)
	g.PrintGraph()
}

type Vertex struct {
	key      int
	vertices []*Vertex
}

type Graph struct {
	vertexs []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if Contain(g.vertexs, k) {
		fmt.Printf("Vertex already exists...")

	} else {
		g.vertexs = append(g.vertexs, &Vertex{key: k})
	}

}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Printf("Error one of the vertext does not exists")

	} else if Contain(fromVertex.vertices, to) {
		fmt.Printf("Existing edge %v-->%v\n", from, to)

	} else {
		fromVertex.vertices = append(fromVertex.vertices, toVertex)
	}

}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertexs {
		if v.key == k {
			return g.vertexs[i]
		}

	}
	return nil
}

func Contain(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}

	}
	return false
}

func (g *Graph) PrintGraph() {
	for _, v := range g.vertexs {
		fmt.Printf("Vertex %v -> ", v.key)
		for _, v := range v.vertices {
			fmt.Printf(" %v ", v.key)

		}
		fmt.Println()

	}
}
