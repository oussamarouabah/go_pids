package pan

import "fmt"

//Greedy algorithm
func (p *Pan) Greedy() {
	p.delta()
	p.unsat()
	p.sort(0, p.N-1)
	for _, v := range p.Delta {
		if delta, index := p.DeltaRef[v[0]].getDelta(), v[0]; delta > 0 {
			for j := 0; j < delta; j++ {
				u := p.getMaxUnsat(p.getUndominatedNeighbors(index))
				p.add(u)
				for _, v := range p.AdjList[u] {
					p.decrimantDelta(v)
				}
			}
		}
	}
	fmt.Println(len(p.DominatingSet), p.N)
}

func (p *Pan) decrimantDelta(index int) {
	if p.DeltaRef[index].getDelta() > 0 {
		p.DeltaRef[index].reference++
		if p.DeltaRef[index].getDelta() == 0 {
			for _, item := range p.AdjList[index] {
				p.Unsat[item].reference++
			}
		}
	}
}
