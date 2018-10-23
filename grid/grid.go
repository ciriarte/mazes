package grid

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strings"

	"github.com/ciriarte/base36"
	"github.com/ciriarte/cell"
)

// Grid represents a Maze Data model
type Grid struct {
	rows, columns int
	cells         [][]cell.Cell
}

// New creates a new Grid
func New(rows, columns int) *Grid {
	g := &Grid{
		rows:    rows,
		columns: columns,
		cells:   make([][]cell.Cell, rows),
	}
	for i := range g.cells {
		g.cells[i] = make([]cell.Cell, columns)
	}
	g.prepareGrid()
	g.prepareCells()
	return g
}

// Cells returns the cells
func (g *Grid) Cells() [][]cell.Cell {
	return g.cells
}

func (g *Grid) prepareGrid() {
	for i := range g.cells {
		for j := range g.cells[i] {
			m := make(map[*cell.Cell]bool)
			g.cells[i][j] = cell.Cell{
				Row:      i,
				Column:   j,
				Distance: math.MaxInt32,
				Links:    m,
			}
		}
	}
}

func (g *Grid) prepareCells() {
	for i := range g.cells {
		for j := range g.cells[i] {
			cell := &g.cells[i][j]
			cell.Distance = 15
			if i-1 >= 0 {
				cell.North = &g.cells[i-1][j]
			}
			if j+1 < g.columns {
				cell.East = &g.cells[i][j+1]
			}
			if j-1 >= 0 {
				cell.West = &g.cells[i][j-1]
			}
			if i+1 < g.rows {
				cell.South = &g.cells[i+1][j]
			}
		}
	}
}

// Size returns the number of cells
func (g *Grid) Size() int {
	return g.rows * g.columns
}

func (g *Grid) String() string {
	var builder strings.Builder
	builder.WriteString("+")
	for range g.cells[0] {
		builder.WriteString("---+")
	}
	builder.WriteString("\n")
	for i := range g.cells {
		var top, bottom strings.Builder
		top.WriteString("|")
		bottom.WriteString("+")
		for j := range g.cells[i] {
			cell := &g.cells[i][j]
			top.WriteString(fmt.Sprintf(" %s ", base36.Encode(uint64(cell.Distance))))
			if cell.IsLinked(cell.East) {
				top.WriteString(" ")
			} else {
				top.WriteString("|")
			}

			if cell.IsLinked(cell.South) {
				bottom.WriteString("   ")
			} else {
				bottom.WriteString("---")
			}
			bottom.WriteString("+")
		}
		top.WriteString("\n")
		bottom.WriteString("\n")
		builder.WriteString(top.String())
		builder.WriteString(bottom.String())
	}

	return builder.String()
}

// SaveAsPNG persists the maze as PNG
func (g *Grid) SaveAsPNG(cellSize, maxDistance int) {
	img := image.NewNRGBA(image.Rect(0, 0, 1+g.columns*cellSize, 1+g.rows*cellSize))

	//clear(img, colour)

	for row := range g.cells {
		for column := range g.cells[row] {
			c := g.cells[row][column]
			colour := color.NRGBA{R: 0x55, G: uint8((maxDistance - c.Distance) * 255 / maxDistance), B: 0x55, A: 255}

			if c.North == nil {
				line(img, column*cellSize, row*cellSize, (1+column)*cellSize, row*cellSize, color.Black)
			}
			if c.West == nil {
				line(img, column*cellSize, row*cellSize, column*cellSize, (1+row)*cellSize, color.Black)
			}

			if !c.IsLinked(c.East) {
				line(img, (1+column)*cellSize, row*cellSize, (1+column)*cellSize, (1+row)*cellSize, color.Black)
			} else {
				line(img, (1+column)*cellSize, row*cellSize+1, (1+column)*cellSize, (1+row)*cellSize-1, colour)
			}

			if !c.IsLinked(c.South) {
				line(img, column*cellSize, (1+row)*cellSize, (1+column)*cellSize, (1+row)*cellSize, color.Black)
			} else {
				line(img, column*cellSize+1, (1+row)*cellSize, (1+column)*cellSize-1, (1+row)*cellSize, colour)
			}
			for x := column*cellSize + 1; x < (1+column)*cellSize; x++ {
				for y := row*cellSize + 1; y < (1+row)*cellSize; y++ {
					img.Set(x, y, colour)
				}
			}
		}
	}

	f, err := os.Create("maze.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func line(img *image.NRGBA, x1, y1, x2, y2 int, colour color.Color) {
	steep := false
	if math.Abs(float64(x1-x2)) < math.Abs(float64(y1-y2)) { // if the line is steep, we transpose the image
		x1, y1 = y1, x1
		x2, y2 = y2, x2
		steep = true
	}
	if x1 > x2 { // make it left−to−right
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	for x := x1; x <= x2; x++ {
		t := float64(x-x1) / float64(x2-x1)
		y := float64(y1)*(1.-t) + (float64(y2) * t)
		if steep {
			img.Set(int(y), x, colour) // if transposed, de−transpose
		} else {
			img.Set(x, int(y), colour)
		}
	}
}

func clear(img *image.NRGBA, background color.Color) {
	maxX := img.Bounds().Size().X
	maxY := img.Bounds().Size().Y
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			img.Set(x, y, background)
		}
	}
}
