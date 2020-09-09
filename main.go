package main

import (
	"fmt"
	"os"

	check "github.com/oussamarouabah/pids_go/PIDS/checkSolution"
	"github.com/oussamarouabah/pids_go/PIDS/graph"
	"github.com/oussamarouabah/pids_go/PIDS/pan"
	"github.com/oussamarouabah/pids_go/PIDS/rei"
	"github.com/oussamarouabah/pids_go/PIDS/wang"
)

func main() {
	g := graph.New(os.Args[1])
	w := wang.New(g)
	w.Greedy()
	ch := check.New(g, w.DominatingSet)
	ch.CheckSolution()
	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("/////////////////////////////////////////////////")

	r := rei.New(g)
	r.Greedy()
	ch = check.New(g, r.DominatingSet)
	ch.CheckSolution()

	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("/////////////////////////////////////////////////")

	p := pan.New(g)
	p.Greedy()
	ch = check.New(g, r.DominatingSet)
	ch.CheckSolution()
}
