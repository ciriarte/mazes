package dijkstra

import (
	"container/heap"
	"math"

	"github.com/ciriarte/mazes/cell"
	"github.com/ciriarte/mazes/grid"
)

// Item a cell with a distance
type Item struct {
	value    *cell.Cell
	distance int
	index    int
}

// MinQueue maintains the heap property for smallest item
type MinQueue []*Item

// Len returns the number of items in MinQueue
func (md MinQueue) Len() int { return len(md) }

// Less compares two items in the priority queue
func (md MinQueue) Less(i, j int) bool {
	if md[i] == nil || md[j] == nil {
		return false
	}
	return md[i].distance < md[j].distance
}

// Swap exchanges the items at index i and j
func (md MinQueue) Swap(i, j int) {
	md[i], md[j] = md[j], md[i]
	md[i].index = i
	md[j].index = j
}

// Pop removes an Item from the heap
func (md *MinQueue) Pop() interface{} {
	old := *md
	n := len(old)
	item := old[n-1]
	*md = old[:n-1]
	item.index = n
	return item
}

// Push adds an Item to the heap
func (md *MinQueue) Push(x interface{}) {
	n := len(*md)
	item := x.(*Item)
	item.index = n
	*md = append(*md, item)
}

func (md *MinQueue) update(item *Item) {
	for _, v := range *md {
		if v.value.Row == item.value.Row &&
			v.value.Column == item.value.Column {
			v.distance = item.distance
			heap.Fix(md, v.index)
			return
		}
	}
}

// Run mutates a Grid object
func Run(g *grid.Grid, source *cell.Cell) (map[*cell.Cell]int, map[*cell.Cell]*cell.Cell) {
	previous := make(map[*cell.Cell]*cell.Cell, g.Size())
	unvisited := make(MinQueue, 0)
	distance := make(map[*cell.Cell]int, g.Size())

	heap.Init(&unvisited)
	source.Distance = 0
	distance[source] = 0
	heap.Push(&unvisited, &Item{
		value:    source,
		distance: 0,
		index:    0,
	})

	cells := g.Cells()
	index := 0
	for row := range cells {
		for column := range cells[row] {
			c := &cells[row][column]
			if c != source {
				heap.Push(&unvisited, &Item{
					value:    c,
					distance: math.MaxInt32,
					index:    index,
				})
			}
			index++
		}
	}

	for len(unvisited) > 0 {
		c := (heap.Pop(&unvisited)).(*Item)
		neighbors := c.value.Neighbors()
		for _, v := range neighbors {
			newDistance := c.distance + 1
			_, ok := distance[v]
			if !ok {
				distance[v] = math.MaxInt32
			}
			if newDistance < distance[v] {
				distance[v] = newDistance
				v.Distance = newDistance
				previous[v] = c.value
				unvisited.update(&Item{
					value:    v,
					distance: newDistance,
					index:    -1,
				})
			}
		}
	}

	return distance, previous
}
