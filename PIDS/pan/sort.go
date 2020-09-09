package pan

func (p *Pan) sort(begin int, end int) {
	if begin < end {
		q := p.partition(begin, end)
		p.sort(begin, q-1)
		p.sort(q+1, end)
	}
}

func (p *Pan) partition(begin int, end int) int {
	x := p.Delta[end][1]
	i := begin - 1
	for j := begin; j < end; j++ {
		if p.Delta[j][1] > x {
			i++
			p.Delta[i][0], p.Delta[j][0] = exchange(p.Delta[i][0], p.Delta[j][0])
			p.Delta[i][1], p.Delta[j][1] = exchange(p.Delta[i][1], p.Delta[j][1])
		}
	}
	p.Delta[end][0], p.Delta[i+1][0] = exchange(p.Delta[end][0], p.Delta[i+1][0])
	p.Delta[end][1], p.Delta[i+1][1] = exchange(p.Delta[end][1], p.Delta[i+1][1])
	return i + 1
}

func exchange(x int, y int) (int, int) {
	return y, x
}
