package main

import (
	"fmt"
	"os"
	"time"

	"github.com/oussamarouabah/pids_go/PIDS/graph"
	"github.com/oussamarouabah/pids_go/PIDS/raei"
)

func main() {
	g := graph.New(os.Args[1])
	size, collapsed := run(g)
	fmt.Println(size, ",", collapsed)
}

func run(g *graph.Graph) (int, int64) {
	w := raei.New(g)
	start := time.Now()
	w.Greedy()
	end := time.Now()
	collapsed := end.Sub(start)
	return len(w.DominatingSet), collapsed.Microseconds()
}
