package heuristcs

//Init befor start the algo
func (h *Heuristics) Init() {

	h.DominatingSet = []int{}

	for i := 0; i < h.N; i++ {
		h.need[i] = (len(h.AdjList[i]) + 1) / 2
	}

	for i := 0; i < h.N; i++ {
		for _, v := range h.AdjList[i] {
			h.utility[i] += h.need[v]
		}
	}

}
