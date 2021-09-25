package graph

import (
	"fmt"
	"funcParty2/arrays"
	"math/rand"
)

type Node struct {
	Label       string
	value       interface{}
	connections []Node
}

func (n Node) HasChild(child *Node) bool {
	for i := 0; i < len(n.connections); i++ {
		if n.connections[i].Label == child.Label {
			return true
		}
	}
	return false
}

type Graph struct {
	nodes []Node
}

func (g *Graph) Gen() {
	numOfNodes := 100
	connections := 68
	preSec := []string{}
	for i := 0; i < numOfNodes; i++ {
		temp := uniqRandSeq(2, preSec)
		preSec = append(preSec, temp)
		node1 := Node{
			Label:       temp,
			value:       nil,
			connections: nil,
		}
		g.nodes = append(g.nodes, node1)
	}

	for i := 0; i < connections; i++ {
		//rand.Seed(time.Now().UnixNano()+654644)
		g.makeConnection(&g.nodes[rand.Intn(len(g.nodes))], &g.nodes[rand.Intn(len(g.nodes))])
	}
}

func (g *Graph) makeConnection(n *Node, n2 *Node) {
	if n.Label == n2.Label {
		return
	}
	if n.HasChild(n2) {
		return
	}
	if n2.HasChild(n) {
		return
	}

	n.connections = append(n.connections, *n2)
	n2.connections = append(n2.connections, *n)
}

func (g *Graph) Show() {
	for index, n := range g.nodes {
		fmt.Println(index, n.Label)
		for _, n2 := range n.connections {
			fmt.Println("\t\t", n2.Label)
		}
	}
}

func (g Graph) BFS(start Node, unvisited []Node) []Node {
	var visited []Node
	var queue arrays.Queue
	queue.Push(start)

	for len(queue.Q) != 0 {
		n := queue.Pop().(Node)
		visited = append(visited, n)
		for i := 0; i < len(unvisited); i++ {
			if n.HasChild(&unvisited[i]) && !isVisited(unvisited[i], visited) {
				queue.Push(unvisited[i])
			}
		}
	}
	return visited

}

func (g *Graph) FindAllConnectives() [][]Node {
	var all [][]Node
	var unvisited []Node = g.nodes
	for len(unvisited) != 0 {
		visited := g.BFS(unvisited[0], unvisited)
		all = append(all, visited)

		for i := 0; i < len(visited); i++ {
			for j := 0; j < len(unvisited); j++ {
				if unvisited[j].Label == visited[i].Label {
					unvisited = remove(unvisited, j)
				}
			}
		}

	}
	return all
}

func isVisited(n Node, visited []Node) bool {
	for i := 0; i < len(visited); i++ {
		if n.Label == visited[i].Label {
			return true
		}
	}
	return false
}

func remove(s []Node, i int) []Node {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func uniqRandSeq(n int, preSec []string) string {
	//rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	for isThere(string(b), preSec) {
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
	}
	return string(b)

}

func isThere(s string, sec []string) bool {
	for i := 0; i < len(sec); i++ {
		if s == sec[i] {
			return true
		}
	}
	return false
}
