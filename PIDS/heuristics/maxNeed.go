package heuristcs

//MaxNeed heuristic
func (h *Heuristics) MaxNeed() {

	needs := h.getSumNeed()

	for needs > 0 {

		u := h.GetMaxNeed(h.getDominatedSet())

		if u == -1 {
			continue
		}

		h.DominatingSet = append(h.DominatingSet, u)

		needs -= h.need[u]
		h.colors[u] = 1

		for _, v := range h.AdjList[u] {

			h.reference[v] ++

			if h.reference[v] >= h.need[v] {
				h.colors[v] = 1
			}

			if h.colors[v] == 0 {
				needs--

				for _, z := range h.AdjList[v] {

					if h.utility[z] > 0 {

						h.utility[z]--

					}

				}
			}
			

			if h.utility[v] > 0 {

				h.utility[v] -= h.need[u]

			}

		}

	}

}
