package heuristcs

//MaxNeed heuristic
func (h *Heuristics) MaxNeed() {
	needs := h.getSumNeed()
	for needs > 0 {
		u := h.GetMaxNeed(h.getDominatedSet())
		if u == -1 {
			continue;
		}
		h.DominatingSet = append(h.DominatingSet,u)
		needs -= h.need[u]
		for _,v := range h.g.AdjList[u] {
			if h.need[v]  > 0 {
				h.need[v] --
				needs --
				for _,z := range h.g.AdjList[v] {
					if h.utility[z] > 0 {
						h.utility[z] --
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