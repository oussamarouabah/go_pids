package heuristcs

import "fmt"

//ShowSolution show solution
func (h *Heuristics) ShowSolution() {

	fmt.Println("Solution size :: ", len(h.DominatingSet))

	for _, v := range h.DominatingSet {
		fmt.Print(v+1, " ")
	}
	fmt.Print("\n")
}

//CheckSolution print needs
func (h *Heuristics) CheckSolution() {
	for i := 0; i < h.N; i++ {
		if h.needReference[i] < h.need[i] || h.colors[i] == 0 {
			fmt.Println("colors(", i+1, ") = ", h.colors[i])
		}
	}
}
