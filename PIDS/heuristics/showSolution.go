package heuristcs

import "fmt"

//ShowSolution show solution
func (h *Heuristics) ShowSolution() {
	fmt.Println("Solution size :: ", len(h.DominatingSet))
	for _, v := range h.DominatingSet {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}
