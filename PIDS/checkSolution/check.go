package check

import (
	"fmt"

	"github.com/oussamarouabah/pids_go/PIDS/graph"
)

type dominatingSet []int

func (d dominatingSet) find(x int) bool {
	for _, v := range d {
		if v == x {
			return true
		}
	}
	return false
}

type Check struct {
	*graph.Graph
	Solution dominatingSet
	need     []int
	ref      []int
}

func New(g *graph.Graph, solution []int) *Check {
	return &Check{
		Graph:    g,
		Solution: solution,
		need:     make([]int, g.N),
		ref:      make([]int, g.N),
	}
}

func (c *Check) calculateNeed() {
	for i := 0; i < c.N; i++ {
		c.need[i] = (len(c.AdjList[i]) + 1) / 2
	}
}

func (c *Check) CheckSolution() {
	c.calculateNeed()
	for _, vertex := range c.Solution {
		for _, v := range c.AdjList[vertex] {
			c.ref[v]++
		}
	}
	c.check()
}

func (c *Check) check() {
	for k, v := range c.need {
		if c.ref[k] < v && !c.Solution.find(k) {
			fmt.Println("error reference of", k, "smaller than need")
		}
	}
}
