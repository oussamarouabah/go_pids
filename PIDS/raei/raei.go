package raei

import (
	"github.com/oussamarouabah/pids_go/PIDS/graph"
)

type need struct {
	fixed     int
	reference int
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

//Rei struct holds the needed data
type Raei struct {
	*graph.Graph
	DominatingSet  dominatingSet
	NeedDegree     []need
	CoverDegree    []cover
	UndominatedSet []int
}

//New function return an instance of the data
func New(g *graph.Graph) *Raei {
	r := &Raei{
		Graph:          g,
		DominatingSet:  make([]int, 0),
		NeedDegree:     make([]need, g.N),
		CoverDegree:    make([]cover, g.N),
		UndominatedSet: make([]int, g.N),
	}
	for index := 0; index < g.N; index++ {
		r.UndominatedSet[index] = index
	}
	return r
}

func (r *Raei) add(vertex int) {
	r.DominatingSet = append(r.DominatingSet, vertex)
	r.remove(vertex)
}

func (r *Raei) remove(vertex int) {
	for k, v := range r.UndominatedSet {
		if v == vertex {
			r.UndominatedSet = append(r.UndominatedSet[:k], r.UndominatedSet[k+1:]...)
			return
		}
	}
}

func (r *Raei) need() {
	for i := 0; i < r.N; i++ {
		r.NeedDegree[i].fixed = (len(r.AdjList[i]) + 1) / 2
	}
}

func (r *Raei) coverDegree() {
	for i := 0; i < r.N; i++ {
		for _, v := range r.AdjList[i] {
			r.CoverDegree[i].fixed += r.NeedDegree[v].fixed
		}
	}
}

func (r *Raei) init() {
	r.need()
	r.coverDegree()
}

func (r *Raei) getSumNeed() (result int) {
	for _, v := range r.NeedDegree {
		result += v.fixed
	}
	return
}

func (r *Raei) getMaxCover() int {
	max := r.CoverDegree[r.UndominatedSet[0]].getCover()
	vertex := r.UndominatedSet[0]
	for _, value := range r.UndominatedSet {
		if r.CoverDegree[value].getCover() > max {
			max, vertex = r.CoverDegree[value].getCover(), value
		}
	}
	return vertex
}

//Greedy method of rei's
func (r *Raei) Greedy() {
	r.init()
	for sum := r.getSumNeed(); sum > 0; {
		vertex := r.getMaxCover()
		r.add(vertex)
		sum -= r.NeedDegree[vertex].getNeed()
		for _, v := range r.AdjList[vertex] {
			r.CoverDegree[v].reference += r.NeedDegree[vertex].getNeed()
			if r.NeedDegree[v].getNeed() > 0 {
				r.NeedDegree[v].reference++
				sum--
				for _, u := range r.AdjList[v] {
					r.CoverDegree[u].reference++
				}
			}
		}
		r.NeedDegree[vertex].reference = r.NeedDegree[vertex].fixed
	}
}