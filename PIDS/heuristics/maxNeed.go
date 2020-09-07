package heuristcs

//MaxNeed heuristic
func (h *Heuristics) MaxNeed() {
	needs := 0
	for needs < h.N {
		u := h.GetMaxNeed(h.getUnDominatedSet())
		if u == -1 {
			continue
		}
		h.DominatingSet = append(h.DominatingSet, u)
		if h.colors[u] == 0 {
			needs++
			h.colors[u] = 1
		}
		for _, v := range h.AdjList[u] {
			if h.needReference[v] < h.need[v] {
				h.needReference[v]++
				h.utilityReference[v] += (h.need[u] - h.needReference[u])
				if h.needReference[v] == h.need[v] && h.colors[v] == 0 {
					h.colors[v] = 1
					needs++
				}
				for _, value := range h.AdjList[v] {
					if h.utilityReference[value] < h.utility[value] {
						h.utilityReference[value]++
					}
				}
			}
		}
	}
}
