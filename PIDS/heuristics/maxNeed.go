package heuristcs

import "fmt"

//MaxNeed heuristic
func (h *Heuristics) MaxNeed() {

	needs := 0

	for needs < h.N {

		u := h.GetMaxNeed(h.getDominatedSet())
		h.DominatingSet = append(h.DominatingSet, u)
		needs++
		h.colors[u] = 1
		h.reference[u] = h.need[u]
		fmt.Println("dominanat :: ", u+1)
		for _, v := range h.AdjList[u] {
			h.reference[v]++
			if h.reference[v] >= h.need[v] && h.colors[v] == 0 {
				h.colors[v] = 1
				needs++
			}
			fmt.Println("ref(", v+1, ") = ", h.reference[v], "color(", v+1, ") = ", h.colors[v], "need(", v+1, ") = ", h.need[v])
		}

	}

}
