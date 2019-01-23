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
	grid     []*tile
	entrance *tile
}

func newBoard(x, y int) *board {
	grid := make([]*tile, 0, x*y)
	for i := 0; i < x*y; i++ {
		grid = append(grid, newTile(i, x, y))
	}

	return &board{
		x:        x,
		y:        y,
		grid:     grid,
		entrance: random(grid, wallVertical, wallHorizontal),
	}
}

func (b *board) print() {
	for i, t := range b.grid {
		fmt.Print(t.value)
		if i%b.x == b.x-1 {
			fmt.Println()
		}
	}
}

func (b *board) next(t *tile) []direction {
	var d []direction

	upI := (t.y-1)*b.x + t.x
	if 0 < upI && b.grid[upI].value == empty {
		d = append(d, up)
	}

	downI := (t.y+1)*b.x + t.x
	if downI < len(b.grid) && b.grid[downI].value == empty {
		d = append(d, down)
	}

	leftI := t.y*b.x + t.x - 1
	if 0 < leftI && b.grid[leftI].value == empty {
		d = append(d, left)
	}

	rightI := t.y*b.x + t.x + 1
	if rightI < len(b.grid) && b.grid[rightI].value == empty {
		d = append(d, right)
	}

	return d
}

func (b *board) step(t *tile, d direction) *tile {
	var i int

	switch d {
	case up:
		i = (t.y-1)*b.x + t.x
	case down:
		i = (t.y+1)*b.x + t.x
	case left:
		i = t.y*b.x + t.x - 1
	case right:
		i = t.y*b.x + t.x + 1
	}

	if 0 < i && i < len(b.grid) {
		return b.grid[i]
	}

	return nil
}

func (b *board) walk(t1, t2 *tile, d direction) bool {
	t := t1

	for t.x != t2.x || t.y != t2.y {
		t = b.step(t, d)
		if t.value == boulder {
			return false
		}
	}

	return true
}

func (b *board) mark(t1, t2 *tile, d direction, m tileType) {
	t := t1
	for t.x != t2.x || t.y != t2.y {
		t = b.step(t, d)
		t.value = m
	}
}

func (b *board) nextSolutionTile(p *tile, d direction, exit bool, invalid []int) (*tile, direction) {
	var v direction
	var i int
	var isWrongTileType bool

	for i == 0 || i == (p.y*b.x+p.x) || isWrongTileType {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		if d == up || d == down {
			if len(invalid) == b.y {
				log.Fatal("failed to build board, try again")
			}

			y := r.Intn(b.y)
			if contains(invalid, y) {
				continue
			}

			i = y*b.x + p.x
			v = up
			if y-p.y > 0 {
				v = down
			}
		} else {
			if len(invalid) == b.x {
				log.Fatal("failed to build board, try again")
			}

			x := r.Intn(b.x)
			if contains(invalid, x) {
				continue
			}

			i = p.y*b.x + x
			v = left
			if x-p.x > 0 {
				v = right
			}
		}

		isWrongTileType = b.grid[i].value != empty
		if exit {
			isWrongTileType = b.grid[i].value != wallHorizontal && b.grid[i].value != wallVertical
		}
	}

	tn := b.step(b.grid[i], v)
	walkable := b.walk(p, b.grid[i], v)
	validNextTile := b.grid[i].value == empty && tn.value == empty
	if walkable && (exit || validNextTile) {
		b.mark(p, b.grid[i], v, path)
		return b.grid[i], v
	}

	return b.nextSolutionTile(p, d, exit, append(invalid, i))
}

func (b *board) solve(n int) []direction {
	b.entrance.value = entrance
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
		tn.value = boulder

		// toggle d axis for next path
		if v == left || v == right {
			prevD = up
		} else {
			prevD = left
		}
	}

	prev, prevD = b.nextSolutionTile(prev, prevD, true, []int{})
	solution = append(solution, prevD)
	prev.value = exit

	return solution
}

func random(grid []*tile, only ...tileType) *tile {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(grid))

	for _, t := range only {
		if t == grid[i].value {
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
