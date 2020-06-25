package heuristcs

import "fmt"

//MaxNeedMaxUtility heuristic
func (h *Heuristics) MaxNeedMaxUtility() {
	needs := 0
	for needs < h.N {
		z := h.GetMaxNeed(h.getUnDominatedSet())
		u := h.GetMaxUtility(h.getUnDominatingNeighbors(z))
		if u == -1 {
			break
		}
		h.DominatingSet = append(h.DominatingSet, u)
		if h.colors[u] == 0 {
			h.colors[u] = 1
			needs++
		}
		h.needReference[u] = h.need[u]
		for _, v := range h.AdjList[u] {
			h.needReference[v]++
			if h.needReference[v] == h.need[v] && h.colors[v] == 0 {
				h.colors[v] = 1
				needs++
				fmt.Println("V ::", v, "needs ::", needs)
			}
		}
	}
}
