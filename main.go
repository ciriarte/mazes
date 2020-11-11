package main

import (
	"fmt"
	"math"

	"github.com/ciriarte/mazes/dijkstra"
	"github.com/ciriarte/mazes/grid"
	"github.com/ciriarte/mazes/sidewinder"
)

func main() {
	g := grid.New(50, 50)
	sidewinder.Run(g)
	cells := g.Cells()
	distance, _ := dijkstra.Run(g, &cells[0][0])

	maxDistance := math.MinInt32
	for k := range distance {
		if maxDistance < distance[k] && distance[k] != math.MaxInt32 {
			maxDistance = distance[k]
		}
	}
	fmt.Printf("Max distance: %d", maxDistance)

	g.SaveAsPNG(10, maxDistance)
}
