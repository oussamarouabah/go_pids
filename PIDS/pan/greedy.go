package pan

import "fmt"

//Greedy algorithm
func (p *Pan) Greedy() {
	p.delta()
	p.unsat()
	p.sort(0, p.N-1)
	for _, v := range p.Delta {
		if v[1] > 0 {
			r := v[1]
			for j := 0; j < r; j++ {
				u := p.getMaxUnsat(p.getUndominatedNeighbors(v[0]))
				p.DominatingSet = append(p.DominatingSet, u)
				for _, v := range p.AdjList[u] {
					p.decrimantDelta(v)
				}
			}
		}
	}
	fmt.Println(len(p.DominatingSet), p.N)
}

func (p *Pan) decrimantDelta(index int) {
	for k, v := range p.Delta {
		if v[0] == index && v[1] > 0 {
			p.Delta[k][1]--
			if v[1] == 0 {
				for _, item := range p.AdjList[v[0]] {
					p.Unsat[item]--
				}
			}
		}
	}
}
