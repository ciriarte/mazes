package binarytree

import (
	"math/rand"
	"time"

	"github.com/ciriarte/mazes/cell"
	"github.com/ciriarte/mazes/grid"
)

// Run mutates a Grid object
func Run(g *grid.Grid) *grid.Grid {
	rand.Seed(time.Now().UTC().UnixNano())
	cells := g.Cells()
	for i := range cells {
		for j := range cells[i] {
			c := &cells[i][j]
			list := make([]*cell.Cell, 0, 4)
			if c.East != nil {
				list = append(list, c.East)
			}
			if c.North != nil {
				list = append(list, c.North)
			}
			if length := len(list); length > 0 {
				northOrEast := rand.Intn(length)
				c.Link(list[northOrEast], true)
			}
		}
	}

	return g
}
