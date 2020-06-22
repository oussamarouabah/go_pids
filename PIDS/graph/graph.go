package graph

type Graph struct {
	N       int
	AdjList [][]int
	MatAdj  [][]int
}

func New(path string) *Graph {
	//TO Do open file and generate the AdjList && MatList
	return &Graph{}
}

//To do
func (g *Graph) ShowAdjList() {

}

//To do
func (g *Graph) ShowMatList() {

}
