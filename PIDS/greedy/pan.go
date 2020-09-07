package greedy

import (
	"github.com/oussamarouabah/pids_go/PIDS/graph"
)

type Pan struct {
	*graph.Graph
	Solution []int
	Delta []int
}
//New create a new instance
func New(g *graph.Graph) *Pan {
	return &Pan{
		Graph:    g,
	}
}

func (p *Pan) delta() {}
func (p *Pan) unsat() {}
func (p *Pan) getMaxDelta() int { return 0}
func (p *Pan) sort()  {}

//Greedy algorithm
func (p *Pan) Greedy() {
	p.delta()
	p.unsat()
	p.sort()
	for i := 0; i < p.N; i++ {
		//v := p.getMaxDelta()
		if p.Delta[i] > 0 {
			r := p.Delta[i]
			for j := 0; j < r ; j++ {

			}
		}
	}
}
