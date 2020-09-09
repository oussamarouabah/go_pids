package pan

type dominatingSet []int

func (d dominatingSet) find(x int) bool {
	for _, v := range d {
		if v == x {
			return true
		}
	}
	return false
}
