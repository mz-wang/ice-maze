package main

type tileType string

const (
	boulder        tileType = "o"
	corner         tileType = "+"
	empty          tileType = " "
	entrance       tileType = "E"
	exit           tileType = "X"
	path           tileType = "."
	wallVertical   tileType = "|"
	wallHorizontal tileType = "-"
)

type tile struct {
	X     int      `json:"x"`
	Y     int      `json:"y"`
	Value tileType `json:"value"`
}

func newTile(i, x, y int) *tile {
	value := empty
	if i == 0 || i == x-1 || i == x*y-x || i == x*y-1 {
		value = corner
	} else if (0 < i && i < x) || (x*y-x < i && i < x*y-1) {
		value = wallHorizontal
	} else if i%x == 0 || i%x == x-1 {
		value = wallVertical
	}

	return &tile{
		X:     i % x,
		Y:     i / x,
		Value: value,
	}
}
