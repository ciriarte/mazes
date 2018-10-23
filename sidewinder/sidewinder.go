package sidewinder

import (
	"math/rand"
	"time"

	"github.com/ciriarte/cell"
	"github.com/ciriarte/grid"
)

// Run mutates a Grid object
func Run(g *grid.Grid) *grid.Grid {
	rand.Seed(time.Now().UTC().UnixNano())
	run := make([]*cell.Cell, 0)
	cells := g.Cells()
	for i := range cells {
		for j := range cells[i] {
			c := &cells[i][j]
			run = append(run, c)

			atEasterBoundary := c.East == nil
			atNorthernBoundary := c.North == nil

			shouldCloseOut := atEasterBoundary ||
				(!atNorthernBoundary && rand.Intn(2) == 0)

			if shouldCloseOut {
				member := run[rand.Intn(len(run))]
				if member.North != nil {
					member.Link(member.North, true)
				}
				run = make([]*cell.Cell, 0)
			} else {
				c.Link(c.East, true)
			}
		}
	}

	return g
}
