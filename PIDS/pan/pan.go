package pan

import (
	"github.com/oussamarouabah/pids_go/PIDS/graph"
)

//Pan is a struct
type Pan struct {
	*graph.Graph
	DominatingSet dominatingSet
	Delta         [][2]int
	Unsat         []int
	DeltaRef      []int
}

//New create a new instance
func New(g *graph.Graph) *Pan {
	return &Pan{
		Graph:    g,
		Delta:    make([][2]int, g.N),
		Unsat:    make([]int, g.N),
		DeltaRef: make([]int, g.N),
	}
}

func (p *Pan) delta() {
	for i := 0; i < p.N; i++ {
		p.Delta[i][0] = i
		p.Delta[i][1] = (len(p.AdjList[i]) + 1) / 2
	}
}

func (p *Pan) unsat() {
	for i := 0; i < p.N; i++ {
		p.Unsat[i] = len(p.AdjList[i])
	}
}

func (p *Pan) getUndominatedNeighbors(index int) []int {
	result := []int{}
	for _, v := range p.AdjList[index] {
		if !p.DominatingSet.find(v) {
			result = append(result, v)
		}
	}
	return result
}
