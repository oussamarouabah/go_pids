package main

import (
	"fmt"
	"os"
	"time"

	"github.com/oussamarouabah/pids_go/PIDS/graph"

	heuristcs "github.com/oussamarouabah/pids_go/PIDS/heuristics"
)

func main() {
	g := graph.New(os.Args[1])
	//g.ShowAdjList()
	h := heuristcs.New(g)
	// Max needd work fine
	h.Init()
	begin := time.Now().Nanosecond()
	h.MaxNeed()
	end := time.Now().Nanosecond()
	fmt.Println("MaxNeed algo finish in ", end-begin, " ns solution size :: ", len(h.DominatingSet))
	// h.ShowSolution()

	// to do in init calculate the utility
	h.Init()
	begin = time.Now().Nanosecond()
	h.MaxUtility()
	end = time.Now().Nanosecond()
	fmt.Println("MaxUtility algo finish in ", end-begin , " ns solution size :: ", len(h.DominatingSet))
	//h.ShowSolution()

	h.Init()
	begin = time.Now().Nanosecond()
	h.MaxNeedMaxUtility()
	end = time.Now().Nanosecond()
	fmt.Println("MaxNeedMaxUtility algo finish in ", end-begin, " ns solution size :: ", len(h.DominatingSet))
	// h.ShowSolution()

}
