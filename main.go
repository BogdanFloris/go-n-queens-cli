package main

import (
	"flag"
)

// N flag: the number of queens and the size of the board
var N = flag.Int("N", 4, "the number of queens and the size of the board")

func main() {
	flag.Parse()
}
