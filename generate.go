package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	b := generate(40, 20, 50)
	print(b)
}

type TileType string

const (
	WallV   TileType = "|"
	WallH   TileType = "-"
	Corner  TileType = "+"
	Boulder TileType = "o"
	Empty   TileType = " "

	Entrance TileType = "E"
	Exit     TileType = "X"
)

type Tile struct {
	X       int      `json:"x"`
	Y       int      `json:"y"`
	Value   TileType `json:"value"`
	Visited bool
}

type Board struct {
	X        int
	Y        int
	Grid     []Tile
	Entrance Tile
	Exit     Tile
}

func generate(x, y, n int) Board {
	grid := make([]Tile, 0, x*y)

	for i := 0; i < x*y; i++ {
		grid = append(grid, Tile{
			X:     i % x,
			Y:     i / x,
			Value: getTileType(i, x, y),
		})
	}

	ent := random(x, y, true)
	ext := random(x, y, true)
	for abs(ext.X-ent.X) <= 1 && abs(ext.Y-ent.Y) <= 1 {
		ext = random(x, y, true)
	}

	grid[ent.Y*x+ent.X].Value = Entrance
	grid[ext.Y*x+ext.X].Value = Exit

	boulders := 0
	for boulders < n {
		b := random(x, y, false)
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

func getTileType(i, x, y int) TileType {
	if i == 0 || i == x-1 || i == x*y-x || i == x*y-1 {
		return Corner
	} else if (0 < i && i < x) || (x*y-x < i && i < x*y-1) {
		return WallH
	} else if i%x == 0 || i%x == x-1 {
		return WallV
	}
	return Empty
}

func random(x, y int, wall bool) Tile {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(x*y - 1)

	for {
		tileType := getTileType(i, x, y)

		if wall && (tileType == WallH || tileType == WallV) {
			break
		}

		if !wall && tileType == Empty {
			break
		}

		r = rand.New(rand.NewSource(time.Now().UnixNano()))
		i = r.Intn(x*y - 1)
	}

	return Tile{
		X: i % x,
		Y: i / x,
	}
}

func print(board Board) {
	for i, t := range board.Grid {
		fmt.Print(t.Value)
		if i%board.X == board.X-1 {
			fmt.Println()
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
