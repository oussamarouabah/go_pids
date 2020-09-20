package main

import (
	"fmt"
	"os"
	"time"

	"github.com/oussamarouabah/pids_go/PIDS/graph"
	"github.com/oussamarouabah/pids_go/PIDS/pan"
)

func main() {
	var times int64 = 0
	sizes := 0
	g := graph.New(os.Args[1])
	for i := 0; i < 15; i++ {
		size, collapsed := run(g)
		times += collapsed
		sizes += size
	}
	fmt.Println(sizes/15, ",", times/15)
}

func run(g *graph.Graph) (int, int64) {
	w := pan.New(g)
	start := time.Now()
	w.Greedy()
	end := time.Now()
	collapsed := end.Sub(start)
	return len(w.DominatingSet), collapsed.Microseconds()
}

// w := wang.New(g)
// w.Greedy()
//ch := check.New(g, w.DominatingSet)
//ch.CheckSolution()

// r := rei.New(g)
// r.Greedy()

// p := pan.New(g)
// p.Greedy()

// m := mai.New(g)
// m.Greedy()
