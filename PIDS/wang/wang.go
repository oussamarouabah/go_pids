package wang

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

type Wang struct {
	*graph.Graph
	DominatingSet  dominatingSet
	UndominatedSet []int
	NeedDegree     []need
}

//New function return an instance of the data
func New(g *graph.Graph) *Wang {
	w := &Wang{
		Graph:          g,
		DominatingSet:  make([]int, 0),
		UndominatedSet: make([]int, g.N),
		NeedDegree:     make([]need, g.N),
	}
	for index := 0; index < g.N; index++ {
		w.UndominatedSet[index] = index
	}
	return w
}

func (w *Wang) add(vertex int) {
	w.DominatingSet = append(w.DominatingSet, vertex)
	w.remove(vertex)
}

func (w *Wang) remove(vertex int) {
	for k, v := range w.UndominatedSet {
		if v == vertex {
			w.UndominatedSet = append(w.UndominatedSet[:k], w.UndominatedSet[k+1:]...)
			return
		}
	}
}

func (w *Wang) need() {
	for i := 0; i < w.N; i++ {
		w.NeedDegree[i].fixed = (len(w.AdjList[i]) + 1) / 2
	}
}

func (w *Wang) getSumNeed() (result int) {
	for _, value := range w.NeedDegree {
		result += value.fixed
	}
	return
}

func (w *Wang) getMaxF() (vertex int) {
	need, vertex := w.NeedDegree[w.UndominatedSet[0]].getMin(), w.UndominatedSet[0]
	for _, v := range w.UndominatedSet {
		if w.NeedDegree[v].getMin() > need {
			need, vertex = w.NeedDegree[vertex].getMin(), v
		}
	}
	return
}

func (w *Wang) showdata() {
	for i := 0; i < w.N; i++ {
		fmt.Println("need(", i, ") :", w.NeedDegree[i], ":", w.NeedDegree[i].getNeed())
	}
}

func (w *Wang) Greedy() {
	w.need()
	h := w.getSumNeed()
	f := 0
	for f < h {
		vertex := w.getMaxF()
		w.add(vertex)
		f += w.NeedDegree[vertex].getNeed()
		w.NeedDegree[vertex].reference = w.NeedDegree[vertex].fixed
		for _, v := range w.AdjList[vertex] {
			if w.NeedDegree[v].getNeed() > 0 {
				w.NeedDegree[v].reference++
				f++
			}
		}
	}

	fmt.Println(w.DominatingSet, len(w.DominatingSet))
	w.showdata()
}
