package maei

import (
	"math/rand"

	"github.com/oussamarouabah/pids_go/PIDS/graph"
)

type need struct {
	fixed     int
	reference int
}

func (n *need) getMin() int {
	if n.reference < n.fixed {
		return n.reference
	}
	return n.fixed
}

func (n *need) getNeed() int {
	return n.fixed - n.reference
}

type cover struct {
	fixed     int
	reference int
}

func (c *cover) getCover() int {
	return c.fixed - c.reference
}

//Wang structure
type Maei struct {
	*graph.Graph
	DominatingSet  dominatingSet
	UndominatedSet []int
	NeedDegree     []need
	CoverDegree    []cover
}

//New function return an instance of the data
func New(g *graph.Graph) *Maei {
	m := &Maei{
		Graph:          g,
		DominatingSet:  make([]int, 0),
		UndominatedSet: make([]int, g.N),
		NeedDegree:     make([]need, g.N),
		CoverDegree:    make([]cover, g.N),
	}
	for index := 0; index < g.N; index++ {
		m.UndominatedSet[index] = index
	}
	return m
}

func (m *Maei) add(vertex int) {
	m.DominatingSet = append(m.DominatingSet, vertex)
	m.remove(vertex)
}

func (m *Maei) remove(vertex int) {
	for k, v := range m.UndominatedSet {
		if v == vertex {
			m.UndominatedSet = append(m.UndominatedSet[:k], m.UndominatedSet[k+1:]...)
			return
		}
	}
}

func (m *Maei) need() {
	for i := 0; i < m.N; i++ {
		m.NeedDegree[i].fixed = (len(m.AdjList[i]) + 1) / 2
	}
}

func (m *Maei) getSumNeed() (result int) {
	for _, value := range m.NeedDegree {
		result += value.fixed
	}
	return
}

func (m *Maei) coverDegree() {
	for i := 0; i < m.N; i++ {
		for _, v := range m.AdjList[i] {
			m.CoverDegree[i].fixed += m.NeedDegree[v].fixed
		}
	}
}

func (m *Maei) getMaxF() (vertex int) {
	max, vertex := m.NeedDegree[m.UndominatedSet[0]].getMin(), m.UndominatedSet[0]
	for _, v := range m.UndominatedSet {
		if tmax := m.NeedDegree[v].getMin(); tmax > max {
			max, vertex = tmax, v
		}
	}
	result := make([]int, 0)
	for _, v := range m.UndominatedSet {
		if tmax := m.NeedDegree[v].getMin(); tmax == max {
			result = append(result, v)
		}
	}
	if len(result) > 1 {
		vertex = m.getMaxCover(result)
	}
	return
}

func (m *Maei) getMaxCover(list []int) int {
	max := m.CoverDegree[list[0]].getCover()
	for _, value := range list {
		if tmax := m.CoverDegree[value].getCover(); tmax > max {
			max = tmax
		}
	}

	result := make([]int, 0)
	for _, value := range list {
		if m.CoverDegree[value].getCover() == max {
			result = append(result, value)
		}
	}
	return result[rand.Intn(len(result))]
}

//Greedy the actual greedy method of wang.
func (m *Maei) Greedy() {
	m.need()
	m.coverDegree()
	h := m.getSumNeed()
	f := 0
	for f < h {
		vertex := m.getMaxF()
		m.add(vertex)
		f += m.NeedDegree[vertex].getNeed()
		m.NeedDegree[vertex].reference = m.NeedDegree[vertex].fixed
		for _, v := range m.AdjList[vertex] {
			if m.NeedDegree[v].getNeed() > 0 {
				m.NeedDegree[v].reference++
				f++
				for _, u := range m.AdjList[v] {
					m.CoverDegree[u].reference++
				}
			}
		}
	}
}
