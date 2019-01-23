package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func usage() {
	s := `./gen <x> <y> <n>
	x - width of maze
	y - height of maze
	n - number of randomly placed boulders in maze`
	fmt.Println(s)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 4 {
		flag.Usage()
	}

	var x, y, n int
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

	// m, err = strconv.Atoi(flag.Arg(3))
	// if err != nil {
	// 	flag.Usage()
	// }

	b := newBoard(x, y)
	s := b.solve(n)
	b.print()
	fmt.Println(s)
}
