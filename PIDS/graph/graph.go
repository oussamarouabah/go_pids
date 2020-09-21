package graph

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

//Graph structure
type Graph struct {
	N       int
	M       int
	AdjList [][]int
}

//New generete a graph
func New(path string) *Graph {
	//TO Do open file and generate the AdjList && MatList
	dataFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("cannot open file ..")
	}

	parts := regexp.MustCompile("([0-9]+)").FindAllString(string(dataFile), -1)

	N, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	M, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	//Declare the graph instance.
	graph := Graph{
		N:       N,
		M:       M,
		AdjList: make([][]int, N),
	}

	parts = parts[2:]

	for i := 0; i < len(parts); i += 2 {

		v, err := strconv.Atoi(parts[i])
		if err != nil {
			log.Fatal(err)
		}

		z, err := strconv.Atoi(parts[i+1])
		if err != nil {
			log.Fatal(err)
		}

		graph.AdjList[v-1] = append(graph.AdjList[v-1], z-1)
		graph.AdjList[z-1] = append(graph.AdjList[z-1], v-1)
	}
	return &graph
}

//ShowAdjList show list
func (g *Graph) ShowAdjList() {

	for k, v := range g.AdjList {

		fmt.Print(k, " :: ")
		for _, z := range v {
			fmt.Printf("%d  ", z)
		}
		fmt.Println(" ")

	}
}
