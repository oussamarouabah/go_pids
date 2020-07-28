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

func (h *Heuristics) show() {
	for i := 0; i < h.N; i++ {
		fmt.Println(i+1, ":", "Need", h.need[i] - h.needReference[i], "Utility", h.utility[i] - h.utilityReference[i])
	}
}

//CheckSolution print needs
func (h *Heuristics) CheckSolution() {
	references := make([]int, h.N)
	for _, v := range h.DominatingSet {
		for _, v := range h.AdjList[v] {
			references[v]++
		}
	}
	for k, v := range references {
		if h.need[k] > v {
			fmt.Println("error need bigger than reference", k+1)
		}
		if h.needReference[k] != v {
			fmt.Println("error reference uneqaul at:", k+1, "new ref :", v, "old :", h.needReference[k])
		}
	}
}
