package main

import (
	"fmt"
	"log"
	"os"

	"github.com/oussamarouabah/pids_go/PIDS/graph"

	heuristcs "github.com/oussamarouabah/pids_go/PIDS/heuristics"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("to call this program give it this params :")
		fmt.Println("1. algorithm name which it can be :: ")
		fmt.Println("1.1 maxneed.", "\n1.2 maxutility.", "\n1.3 maxneedmaxutility.")
		fmt.Println("2. Path of the file")
		log.Fatal("read the guide above.")
	}
	switch os.Args[1] {
	case "maxneed":
		g := graph.New(os.Args[2])
		h := heuristcs.New(g)
		h.Init()
		h.MaxNeed()
		fmt.Println("MaxNeed Heuristic :: ")
		h.ShowSolution()
		h.CheckSolution()
	case "maxutility":
		g := graph.New(os.Args[2])
		h := heuristcs.New(g)
		h.Init()
		fmt.Println("MaxUtility Heuristic :: ")
		h.MaxUtility()
		h.ShowSolution()
		h.CheckSolution()
	case "maxneedmaxutility":
		g := graph.New(os.Args[2])
		h := heuristcs.New(g)
		h.Init()
		fmt.Println("MaxNeedMaxUtility Heuristic :: ")
		h.MaxNeedMaxUtility()
		h.ShowSolution()
		h.CheckSolution()
	case "minneedmaxutility":
		g := graph.New(os.Args[2])
		h := heuristcs.New(g)
		h.Init()
		fmt.Println("MaxNeedMaxUtility Heuristic :: ")
		h.MinNeedMaxUtility()
		h.ShowSolution()
		h.CheckSolution()
	default:
		fmt.Println("to call this program give it this params :")
		fmt.Println("1. algorithm name which it can be :: ")
		fmt.Println("1.1 maxneed.", "\n1.2 maxutility.", "\n1.3 maxneedmaxutility.")
		fmt.Println("2. Path of the file")
	}
}
