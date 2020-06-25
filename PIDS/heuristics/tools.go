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
	result := false
	for _, v := range d {
		if v == x {
			result = true
			break
		}
	}
	return result
}

//New init the heuristic
func New(g *graph.Graph) *Heuristics {

	heuristic := Heuristics{
		g,
		make(dominatingSet, g.N),
		make([]int, g.N),
		make([]int, g.N),
		make([]int, g.N),
		make([]int, g.N),
		make([]int, g.N),
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
		if h.colors[i] != 0 {
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
		dominated = append(dominated, i)
	}
	return dominated
}

//GetMaxNeed heuristic
func (h *Heuristics) GetMaxNeed(dominated []int) int {

	max := -1
	for _, v := range dominated {
		// if h.need[v]-h.reference[v] > max {
		// 	max = h.need[v] - h.reference[v]
		// }
		if h.need[v] > max {
			max = h.need[v]
		}
	}

	result := []int{}
	for _, v := range dominated {
		// if h.need[v]-h.reference[v] == max {
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

	max := h.utility[dominated[0]]
	for _, v := range dominated {
		if h.utility[v] > max {
			max = h.utility[v]
		}
	}

	result := []int{}
	for _, v := range dominated {
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
		if h.need[v] < min && h.need[v] > 0 {
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
