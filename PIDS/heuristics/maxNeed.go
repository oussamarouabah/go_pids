package heuristcs

//MaxNeed heuristic
func (h *Heuristics) MaxNeed() {
	needs := 0
	for needs < h.N {
		u := h.GetMaxNeed(h.getUnDominatedSet())
		h.DominatingSet = append(h.DominatingSet, u)
		needs++
		h.colors[u] = 1
		h.needReference[u] = h.need[u]
		for _, v := range h.AdjList[u] {
			h.needReference[v]++
			if h.needReference[v] == h.need[v] && h.colors[v] == 0 {
				h.colors[v] = 1
				needs++
			}
		}
	}
}
