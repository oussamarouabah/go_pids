package main

import (
	"os"

	"github.com/oussamarouabah/pids_go/PIDS/graph"
	"github.com/oussamarouabah/pids_go/PIDS/rei"
)

func main() {
	g := graph.New(os.Args[1])
	r := rei.New(g)
	r.Greedy()
}
