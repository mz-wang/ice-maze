package main

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
