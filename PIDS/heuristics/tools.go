package heuristcs

import (
	"math"
	"math/rand"
	"reflect"

	"github.com/oussamarouabah/pids_go/PIDS/graph"
	"golang.org/x/text/number"
)

//Heuristics data
type Heuristics struct {
	g		graph.Graph
	DominatingSet dominatingSet
	need []int
	utility []int
}

type dominatingSet []int;

func (d *dominatingSet) find(x int) bool {
	result := false 
	for k,v := range d {
		if x == d {
			result = true
		}
	}
	return result
}

//New init the heuristic 
func New(g *graph.Graph) (*Heuristics) {
	heuristic := Heuristics{
		g: g,
		need : make([]int,g.N),
		utility: make([]int,g.N)
	}
	return &heuristic
}

func (h *Heuristics) getSumNeed() int {
	sum := 0
	for k,v := range h.need {
		sum += v
	}
	return sum
}

func (h *Heuristics) getDominatedSet() []int {
	dominated := make([]int)
	
	for i := 0; i < h.g.N; i++ {
		 if h.DominatingSet.find(i) {
			 continue;
		 }
		 dominated = append(dominated,i)
	}
	return dominated
}



func (h *Heuristics) getDominatedNeighbors(x int) []int {
	data := h.g.AdjList[x]
	dominated := make([]int)
	
	for k,i := range data {
		 if h.DominatingSet.find(i) {
			 continue;
		 }
		 dominated = append(dominated,i)
	}
	return dominated
}

//GetMAxNeed heuristic
func (h *Heuristics) GetMAxNeed(dominated []int) int {
	max := -1
	for k,v := range dominated {
		if h.need[v] > max {
			max = h.need[v]
		}
	}
	result := make([]int)
	for k,v := range dominated {
		if h.need[v] == max {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return -1
	}
	
	return result[rand.Intn(len(result))]
}

//GetMAxNeed heuristic
func (h *Heuristics) GetMaxUtility(dominated []int) int {
	max := -1
	for k,v := range dominated {
		if h.utility[v] > max {
			max = h.utility[v]
		}
	}

	result := make([]int)
	for k,v := range dominated {
		if h.utility[v] == max {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return -1
	}
	
	return result[rand.Intn(len(result))]
}

//GetMAxNeed heuristic
func (h *Heuristics) GetMinNeed(dominated []int) int {
	min := -1
	for k,v := range dominated {
		if h.need[v] < min && h.need[v] > 0 {
			min = h.need[v]
		}
	}
	result := make([]int)
	for k,v := range dominated {
		if h.need[v] == min {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return -1
	}
	
	return result[rand.Intn(len(result))]
}