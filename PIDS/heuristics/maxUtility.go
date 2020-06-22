package heuristcs

func (h *Heuristics) MaxUtility() {
	needs := h.getSumNeed()
	for needs > 0 {
		u := h.GetMaxUtility(h.getDominatedSet())
		if u == -1 {
			continue;
		}
		h.DominatingSet = append(h.DominatingSet,u)
		needs -= h.need[u]
		// h.need[u] = 0 h.utility[u] = 0
		for k,v := range h.g.AdjList[u] {
			if h.need[v]  > 0 {
				h.need[v] -= 1
				needs -= 1
				for k,z := range h.g.AdjList[v] {
					if h.utility[z] > 0 {
						h.utility[z] -= 1
					}
				}
			}
			if h.utility[v] > 0 {
				h.utility[v] -= h.need[u]
			}
		}
		h.need[u] = 0
		h.utility[u] = 0
	}
	
}