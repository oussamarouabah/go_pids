package pan

func (p *Pan) getMaxUnsat(list []int) int {
	max, vertex := p.Unsat[list[0]], list[0]
	for _, v := range list {
		if p.Unsat[v] > max {
			max, vertex = p.Unsat[v], v
		}
	}
	return vertex
}
