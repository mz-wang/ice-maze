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
	NewBoard(x, y, n).Print()
}
