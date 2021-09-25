package main

import (
	"fmt"
	graph "funcParty2/graph"
)

func main() {

	var g graph.Graph
	g.Gen()
	g.Show()
	var x [][]graph.Node = g.FindAllConnectives()
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Print(x[i][j].Label + " ")
		}
		fmt.Println("")
	}
}
