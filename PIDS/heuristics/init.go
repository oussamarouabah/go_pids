package heuristcs

import (
	"fmt"
)

//Init befor start the algo
func (h *Heuristics) Init() {
	h.DominatingSet = []int{}
	for i := 0; i < h.g.N; i++ {
		h.need[i] = (len(h.g.AdjList[i]) + 1) / 2
	}
	for i := 0; i < h.g.N; i++ {
		for _, v := range h.g.AdjList[i] {
			h.utility[i] += h.need[v]
		}
	}
}

//ShowNeeds show needs
func (h *Heuristics) ShowNeeds() {
	for i := 0; i < h.g.N; i++ {
		fmt.Println(i, " :: ", h.need[i])
	}
}
