package heuristcs

import (
	"math/rand"

	"github.com/oussamarouabah/pids_go/PIDS/graph"
)

//Heuristics data
type Heuristics struct {
	*graph.Graph
	DominatingSet    dominatingSet
	need             []int
	utility          []int
	colors           []int
	needReference    []int
	utilityReference []int
}

type dominatingSet []int

func (d dominatingSet) find(x int) bool {
	for _, v := range d {
		if v == x {
			return true
		}
	}
	return false
}

//New init the heuristic
func New(g *graph.Graph) *Heuristics {
	heuristic := Heuristics{
		Graph:            g,
		DominatingSet:    make(dominatingSet, g.N),
		need:             make([]int, g.N),
		utility:          make([]int, g.N),
		colors:           make([]int, g.N),
		needReference:    make([]int, g.N),
		utilityReference: make([]int, g.N),
	}
	return &heuristic
}

func (h *Heuristics) getSumNeed() int {
	sum := 0
	for _, v := range h.need {
		sum += v
	}
	return sum
}

func (h *Heuristics) getUnDominatedSet() []int {
	dominated := []int{}
	for i := 0; i < h.N; i++ {
		// if h.colors[i] != 0 {
		// 	continue
		// }
		if h.DominatingSet.find(i) {
			continue
		}
		dominated = append(dominated, i)
	}
	return dominated
}

func (h *Heuristics) getUnDominatingNeighbors(x int) []int {
	data := h.AdjList[x]
	dominated := []int{}
	for _, i := range data {
		if h.DominatingSet.find(i) {
			continue
		}
		// if h.colors[i] != 0 {
		// 	continue
		// }
		dominated = append(dominated, i)
	}
	return dominated
}

//GetMaxNeed heuristic
func (h *Heuristics) GetMaxNeed(dominated []int) int {
	max := -1
	for _, v := range dominated {
		// if ref := h.need[v] - h.needReference[v]; ref > 0 && ref > max {
		// 	max = ref
		// }
		if h.need[v] > max {
			max = h.need[v]
		}
	}

	result := []int{}
	for _, v := range dominated {
		// if ref := h.need[v] - h.needReference[v]; ref == max {
		// 	result = append(result, v)
		// }
		if h.need[v] == max {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return -1
	}
	return result[rand.Intn(len(result))]
}

//GetMaxUtility heuristic
func (h *Heuristics) GetMaxUtility(dominated []int) int {
	max := -1
	for _, v := range dominated {
		// if util := h.utility[v] - h.utilityReference[v]; util > 0 && util > max {
		// 	max = util
		// }
		if h.utility[v] > max {
			max = h.utility[v]
		}
	}
	result := []int{}
	for _, v := range dominated {
		// if h.utility[v]-h.utilityReference[v] == max {
		// 	result = append(result, v)
		// }
		if h.utility[v] == max {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return -1
	}
	return result[rand.Intn(len(result))]
}

//GetMinNeed heuristic
func (h *Heuristics) GetMinNeed(dominated []int) int {
	min := h.need[dominated[0]]
	for _, v := range dominated {
		if h.need[v] < min {
			min = h.need[v]
		}
	}
	result := []int{}
	for _, v := range dominated {
		if h.need[v] == min {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return -1
	}
	return result[rand.Intn(len(result))]
}
