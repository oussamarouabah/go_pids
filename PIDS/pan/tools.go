package pan

func (p *Pan) getMaxUnsat(list []int) int {
	max, vertex := p.Unsat[list[0]].getUnsat(), list[0]
	for _, v := range list {
		if u := p.Unsat[v].getUnsat(); u > max {
			max, vertex = u, v
		}
	}
	return vertex
}
