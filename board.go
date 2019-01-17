package main

import "fmt"

type Board struct {
	X        int
	Y        int
	Grid     []Tile
	Entrance Tile
	Exit     Tile
}

func NewBoard(x, y, n int) Board {
	grid := make([]Tile, 0, x*y)

	for i := 0; i < x*y; i++ {
		grid = append(grid, Tile{
			X:     i % x,
			Y:     i / x,
			Value: GetTileType(i, x, y),
		})
	}

	ent := Random(x, y, true)
	ext := Random(x, y, true)
	for Abs(ext.X-ent.X) <= 1 && Abs(ext.Y-ent.Y) <= 1 {
		ext = Random(x, y, true)
	}

	grid[ent.Y*x+ent.X].Value = Entrance
	grid[ext.Y*x+ext.X].Value = Exit

	boulders := 0
	for boulders < n {
		b := Random(x, y, false)
		i := b.Y*x + b.X
		if grid[i].Value != Boulder {
			grid[i].Value = Boulder
			boulders++
		}
	}

	return Board{
		X:        x,
		Y:        y,
		Grid:     grid,
		Entrance: grid[ent.Y*x+ent.X],
		Exit:     grid[ext.Y*x+ext.X],
	}
}

func (b Board) Print() {
	for i, t := range b.Grid {
		fmt.Print(t.Value)
		if i%b.X == b.X-1 {
			fmt.Println()
		}
	}
}
