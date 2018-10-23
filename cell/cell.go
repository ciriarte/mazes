package cell

import "fmt"

// Cell represents a individual cell in the Maze
type Cell struct {
	Row, Column, Distance    int
	North, South, East, West *Cell
	Links                    map[*Cell]bool
}

func (c *Cell) String() string {
	return fmt.Sprintf("(%d, %d)", c.Row, c.Column)
}

// Link creates a passage from the current instance to another cell
func (c *Cell) Link(cell *Cell, bidirectional bool) {
	c.Links[cell] = true
	if bidirectional {
		cell.Link(c, false)
	}
}

// Unlink removes a passage from the current instance to another cell
func (c *Cell) Unlink(cell *Cell, bidirectional bool) {
	delete(c.Links, cell)
	if bidirectional {
		cell.Unlink(c, false)
	}
}

// IsLinked checks if there's a linked between this cell and another cell
func (c *Cell) IsLinked(cell *Cell) bool {
	return c.Links[cell]
}

// Neighbors returns adjacent cells
func (c *Cell) Neighbors() []*Cell {
	list := make([]*Cell, 0, 4)
	if c.North != nil && c.IsLinked(c.North) {
		list = append(list, c.North)
	}
	if c.South != nil && c.IsLinked(c.South) {
		list = append(list, c.South)
	}
	if c.East != nil && c.IsLinked(c.East) {
		list = append(list, c.East)
	}
	if c.West != nil && c.IsLinked(c.West) {
		list = append(list, c.West)
	}
	return list
}
