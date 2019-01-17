package main

import (
	"math/rand"
	"time"
)

func GetTileType(i, x, y int) TileType {
	if i == 0 || i == x-1 || i == x*y-x || i == x*y-1 {
		return Corner
	} else if (0 < i && i < x) || (x*y-x < i && i < x*y-1) {
		return WallH
	} else if i%x == 0 || i%x == x-1 {
		return WallV
	}
	return Empty
}

func Random(x, y int, wall bool) Tile {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(x*y - 1)

	for {
		tileType := GetTileType(i, x, y)

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

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
