package main

import (
	"fmt"
	"sort"
)

type Direction string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

type Board struct {
	X         int
	Y         int
	Grid      []*Tile
	Entrance  *Tile
	Exit      *Tile
	Solutions [][]Direction
}

func NewBoard(x, y, n int) *Board {
	grid := make([]*Tile, 0, x*y)

	for i := 0; i < x*y; i++ {
		grid = append(grid, &Tile{
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

	return &Board{
		X:        x,
		Y:        y,
		Grid:     grid,
		Entrance: grid[ent.Y*x+ent.X],
		Exit:     grid[ext.Y*x+ext.X],
	}
}

func (b *Board) Print() {
	for i, t := range b.Grid {
		fmt.Print(t.Value)
		if i%b.X == b.X-1 {
			fmt.Println()
		}
	}
}

func (b *Board) Solve() *Board {
	d := b.Empty(b.Entrance)
	if len(d) > 0 {
		c := b.Next(b.Entrance, d[0])
		t := b.Walk(c, d[0], false)
		b.Traverse(t, []Direction{d[0]})
		sort.Slice(b.Solutions, func(i, j int) bool {
			return len(b.Solutions[i]) < len(b.Solutions[j])
		})
	}
	return b
}

func (b *Board) Empty(t *Tile) []Direction {
	var d []Direction

	up := (t.Y-1)*b.X + t.X
	if 0 < up && b.Grid[up].Value == Empty {
		d = append(d, Up)
	}

	down := (t.Y+1)*b.X + t.X
	if down < len(b.Grid) && b.Grid[down].Value == Empty {
		d = append(d, Down)
	}

	left := t.Y*b.X + t.X - 1
	if 0 < left && b.Grid[left].Value == Empty {
		d = append(d, Left)
	}

	right := t.Y*b.X + t.X + 1
	if right < len(b.Grid) && b.Grid[right].Value == Empty {
		d = append(d, Right)
	}

	return d
}

func (b *Board) Next(t *Tile, d Direction) *Tile {
	var i int

	switch d {
	case Up:
		i = (t.Y-1)*b.X + t.X
	case Down:
		i = (t.Y+1)*b.X + t.X
	case Left:
		i = t.Y*b.X + t.X - 1
	case Right:
		i = t.Y*b.X + t.X + 1
	}

	if 0 < i && i < len(b.Grid) {
		return b.Grid[i]
	}

	return nil
}

func (b *Board) Walk(t *Tile, d Direction, path bool) *Tile {
	c := t
	if path && t.Value != Entrance {
		t.Value = Path
	}

	for n := b.Next(t, d); n != nil && (n.Value == Empty || n.Value == Path || n.Value == Exit); n = b.Next(n, d) {
		c = n
		if path && n.Value != Entrance && n.Value != Exit {
			n.Value = Path
		}
	}

	return c
}

func (b *Board) Traverse(t *Tile, s []Direction) {
	t.Visited = true
	e := b.Empty(t)

	for _, d := range e {
		w := b.Walk(t, d, false)

		if w == nil || w.Value == Entrance {
			continue
		}

		solution := make([]Direction, len(s))
		copy(solution, s)
		solution = append(solution, d)

		if w.Value == Exit {
			b.Solutions = append(b.Solutions, solution)
		}

		if w.Value == Empty && !w.Visited {
			b.Traverse(w, solution)
		}
	}
}
