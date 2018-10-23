package main

import (
	"fmt"
	"math"

	"github.com/ciriarte/dijkstra"
	"github.com/ciriarte/grid"
	"github.com/ciriarte/sidewinder"
)

func main() {
	g := grid.New(25, 25)
	sidewinder.Run(g)
	cells := g.Cells()
	distance, _ := dijkstra.Run(g, &cells[0][0])

	maxDistance := math.MinInt32
	for k := range distance {
		if maxDistance < distance[k] && distance[k] != math.MaxInt32 {
			maxDistance = distance[k]
		}
	}
	fmt.Printf("%d", maxDistance)

	g.SaveAsPNG(10, maxDistance)
}
