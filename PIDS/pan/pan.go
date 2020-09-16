package pan

import (
	"github.com/oussamarouabah/pids_go/PIDS/graph"
)

//Pan is a struct
type Pan struct {
	*graph.Graph
	DominatingSet  dominatingSet
	Delta          [][2]int
	Unsat          []unsat
	DeltaRef       []delta
	UndominatedSet []int
}

type delta struct {
	fixed, reference int
}

func (d *delta) getDelta() int {
	return d.fixed - d.reference
}

type unsat struct {
	fixed, reference int
}

func (u *unsat) getUnsat() int {
	return u.fixed - u.reference
}

//New create a new instance
func New(g *graph.Graph) *Pan {
	p := &Pan{
		Graph:          g,
		Delta:          make([][2]int, g.N),
		Unsat:          make([]unsat, g.N),
		DeltaRef:       make([]delta, g.N),
		UndominatedSet: make([]int, g.N),
	}
	for i := 0; i < g.N; i++ {
		p.UndominatedSet[i] = i
	}
	return p
}

func (p *Pan) delta() {
	for i := 0; i < p.N; i++ {
		p.Delta[i][0] = i
		p.Delta[i][1] = (len(p.AdjList[i]) + 1) / 2
		p.DeltaRef[i].fixed = p.Delta[i][1]
	}
}

func (p *Pan) unsat() {
	for i := 0; i < p.N; i++ {
		p.Unsat[i].fixed = len(p.AdjList[i])
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

func (p *Pan) add(vertex int) {
	p.DominatingSet = append(p.DominatingSet, vertex)
	p.remove(vertex)
}

func (p *Pan) remove(vertex int) {
	for k, v := range p.UndominatedSet {
		if v == vertex {
			p.UndominatedSet = append(p.UndominatedSet[:k], p.UndominatedSet[k+1:]...)
			return
		}
	}
}
