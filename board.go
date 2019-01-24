package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type direction string

const (
	up    direction = "U"
	down  direction = "D"
	left  direction = "L"
	right direction = "R"
)

type board struct {
	x, y     int
	entrance *tile
	Grid     []*tile     `json:"grid,omitempty"`
	Solution []direction `json:"solution,omitempty"`
}

func newBoard(x, y int) *board {
	grid := make([]*tile, 0, x*y)
	for i := 0; i < x*y; i++ {
		grid = append(grid, newTile(i, x, y))
	}

	return &board{
		x:        x,
		y:        y,
		Grid:     grid,
		entrance: random(grid, wallVertical, wallHorizontal),
	}
}

func (b *board) print() {
	for i, t := range b.Grid {
		fmt.Print(t.Value)
		if i%b.x == b.x-1 {
			fmt.Println()
		}
	}
}

func (b *board) next(t *tile) []direction {
	var d []direction

	upI := (t.Y-1)*b.x + t.X
	if 0 < upI && b.Grid[upI].Value == empty {
		d = append(d, up)
	}

	downI := (t.Y+1)*b.x + t.X
	if downI < len(b.Grid) && b.Grid[downI].Value == empty {
		d = append(d, down)
	}

	leftI := t.Y*b.x + t.X - 1
	if 0 < leftI && b.Grid[leftI].Value == empty {
		d = append(d, left)
	}

	rightI := t.Y*b.x + t.X + 1
	if rightI < len(b.Grid) && b.Grid[rightI].Value == empty {
		d = append(d, right)
	}

	return d
}

func (b *board) step(t *tile, d direction) *tile {
	var i int

	switch d {
	case up:
		i = (t.Y-1)*b.x + t.X
	case down:
		i = (t.Y+1)*b.x + t.X
	case left:
		i = t.Y*b.x + t.X - 1
	case right:
		i = t.Y*b.x + t.X + 1
	}

	if 0 < i && i < len(b.Grid) {
		return b.Grid[i]
	}

	return nil
}

func (b *board) walk(t1, t2 *tile, d direction) bool {
	t := t1

	for t.X != t2.X || t.Y != t2.Y {
		t = b.step(t, d)
		if t.Value == boulder {
			return false
		}
	}

	return true
}

func (b *board) mark(t1, t2 *tile, d direction, m tileType) {
	t := t1
	for t.X != t2.X || t.Y != t2.Y {
		t = b.step(t, d)
		t.Value = m
	}
}

func (b *board) nextSolutionTile(p *tile, d direction, exit bool, invalid []int) (*tile, direction) {
	var v direction
	var i int
	var isWrongTileType bool

	for i == 0 || i == (p.Y*b.x+p.X) || isWrongTileType {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		if d == up || d == down {
			if len(invalid) == b.y {
				panic("failed to build board")
			}

			y := r.Intn(b.y)
			if contains(invalid, y) {
				continue
			}

			i = y*b.x + p.X
			v = up
			if y-p.Y > 0 {
				v = down
			}
		} else {
			if len(invalid) == b.x {
				log.Fatal("failed to build board")
			}

			x := r.Intn(b.x)
			if contains(invalid, x) {
				continue
			}

			i = p.Y*b.x + x
			v = left
			if x-p.X > 0 {
				v = right
			}
		}

		isWrongTileType = b.Grid[i].Value != empty
		if exit {
			isWrongTileType = b.Grid[i].Value != wallHorizontal && b.Grid[i].Value != wallVertical
		}
	}

	tn := b.step(b.Grid[i], v)
	walkable := b.walk(p, b.Grid[i], v)
	validNextTile := b.Grid[i].Value == empty && tn.Value == empty
	if walkable && (exit || validNextTile) {
		b.mark(p, b.Grid[i], v, path)
		return b.Grid[i], v
	}

	return b.nextSolutionTile(p, d, exit, append(invalid, i))
}

func (b *board) solve(n int) {
	b.entrance.Value = entrance
	prev := b.entrance
	prevD := b.next(prev)[0]
	solution := []direction{}

	for i := 0; i < n-1; i++ {
		// randomly generate next solution tile along same x- or y-axis
		t, v := b.nextSolutionTile(prev, prevD, false, []int{})
		solution = append(solution, v)
		prev = t

		// set boulder
		tn := b.step(t, v)
		tn.Value = boulder

		// toggle d axis for next path
		if v == left || v == right {
			prevD = up
		} else {
			prevD = left
		}
	}

	prev, prevD = b.nextSolutionTile(prev, prevD, true, []int{})
	solution = append(solution, prevD)
	prev.Value = exit
	b.Solution = solution
}

func random(grid []*tile, only ...tileType) *tile {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(grid))

	for _, t := range only {
		if t == grid[i].Value {
			return grid[i]
		}
	}

	return random(grid, only...)
}

func contains(slice []int, i int) bool {
	for _, x := range slice {
		if i == x {
			return true
		}
	}

	return false
}
