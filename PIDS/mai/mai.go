package mai

import (
	"fmt"

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

type Mai struct {
	*graph.Graph
	DominatingSet  dominatingSet
	UndominatedSet []int
	NeedDegree     []need
	CoverDegree    []cover
}

//New function return an instance of the data
func New(g *graph.Graph) *Mai {
	w := &Mai{
		Graph:          g,
		DominatingSet:  make([]int, 0),
		UndominatedSet: make([]int, g.N),
		NeedDegree:     make([]need, g.N),
		CoverDegree:    make([]cover, g.N),
	}
	for index := 0; index < g.N; index++ {
		w.UndominatedSet[index] = index
	}
	return w
}

func (w *Mai) add(vertex int) {
	w.DominatingSet = append(w.DominatingSet, vertex)
	w.remove(vertex)
}

func (w *Mai) remove(vertex int) {
	for k, v := range w.UndominatedSet {
		if v == vertex {
			w.UndominatedSet = append(w.UndominatedSet[:k], w.UndominatedSet[k+1:]...)
			return
		}
	}
}

func (w *Mai) need() {
	for i := 0; i < w.N; i++ {
		w.NeedDegree[i].fixed = (len(w.AdjList[i]) + 1) / 2
	}
}

func (w *Mai) coverDegree() {
	for i := 0; i < w.N; i++ {
		for _, v := range w.AdjList[i] {
			w.CoverDegree[i].fixed += w.NeedDegree[v].fixed
		}
	}
}

func (w *Mai) getSumNeed() (result int) {
	for _, value := range w.NeedDegree {
		result += value.fixed
	}
	return
}

func (w *Mai) getMaxF() (vertex int) {
	need, vertex := w.NeedDegree[w.UndominatedSet[0]].getMin(), w.UndominatedSet[0]
	for _, v := range w.UndominatedSet {
		if w.NeedDegree[v].getMin() > need {
			need, vertex = w.NeedDegree[vertex].getMin(), v
		}
	}
	solution := []int{}
	for _, v := range w.UndominatedSet {
		if w.NeedDegree[v].getMin() == need {
			solution = append(solution, v)
		}
	}
	if len(solution) > 1 {
		vertex = w.getMaxCover(solution)

	}
	return
}

func (w *Mai) getMaxCover(list []int) int {
	max := w.CoverDegree[list[0]].getCover()
	vertex := list[0]
	for _, value := range list {
		if w.CoverDegree[value].getCover() > max {
			max, vertex = w.CoverDegree[value].getCover(), value
		}
	}
	return vertex
}

func (w *Mai) showdata() {
	for i := 0; i < w.N; i++ {
		fmt.Println("need(", i, ") :", w.NeedDegree[i], ":", w.NeedDegree[i].getNeed())
	}
}

func (w *Mai) Greedy() {
	w.need()
	w.coverDegree()
	h := w.getSumNeed()
	f := 0
	for f < h {
		vertex := w.getMaxF()
		w.add(vertex)
		f += w.NeedDegree[vertex].getNeed()
		for _, v := range w.AdjList[vertex] {
			w.CoverDegree[v].reference += w.NeedDegree[v].getNeed()
			if w.NeedDegree[v].getNeed() > 0 {
				w.NeedDegree[v].reference++
				f++
				for _, u := range w.AdjList[v] {
					w.CoverDegree[u].reference++
				}
			}
		}
		w.NeedDegree[vertex].reference = w.NeedDegree[vertex].fixed
	}

	fmt.Println(w.DominatingSet, len(w.DominatingSet))
	w.showdata()
}
