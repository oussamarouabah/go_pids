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

//ShowColors show the colors
func (h *Heuristics) ShowColors() {

	fmt.Println("colors :: ")
	for k, v := range h.colors {
		if v == 0 {
			fmt.Println("reference[", k+1, "] = ", h.reference[k], " need[", k+1, "] = ", h.need[k])
		} else {
			fmt.Print("Colors[", k+1, "] = ", h.colors[k], "\t")
			neighbors := h.AdjList[k]
			for _, value := range neighbors {
				fmt.Print(value, "\t")
			}
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}

//CheckSolution print needs
func (h *Heuristics) CheckSolution() {
	for i := 0; i < h.N; i++ {
		if h.reference[i] < h.need[i] || h.colors[i] == 0 {
			fmt.Println("colors(", i+1, ") = ", h.colors[i])
		}
	}
}
