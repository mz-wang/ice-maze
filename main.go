package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var x, y, n int

func init() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 3 {
		flag.Usage()
	}

	var err error
	x, err = strconv.Atoi(flag.Arg(0))
	if err != nil {
		flag.Usage()
	}

	y, err = strconv.Atoi(flag.Arg(1))
	if err != nil {
		flag.Usage()
	}

	n, err = strconv.Atoi(flag.Arg(2))
	if err != nil {
		flag.Usage()
	}
}

func usage() {
	s := `./gen <x> <y> <n>
	x - width of maze
	y - height of maze
	n - number of randomly placed boulders in maze`
	fmt.Println(s)
	os.Exit(1)
}

func main() {
	b := NewBoard(x, y, n).Solve()
	for len(b.Solutions) != 1 {
		b = NewBoard(x, y, n).Solve()
	}

	t := b.Entrance
	for _, d := range b.Solutions[0] {
		t = b.Walk(t, d, true)
	}

	b.Print()
	fmt.Println("solution", b.Solutions[0])
}
