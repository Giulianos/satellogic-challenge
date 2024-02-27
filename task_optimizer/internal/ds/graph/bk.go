package graph

import "task_optimizer/internal/ds/set"

func BronKerbosch(r, p, x set.Set[int], graph Graph) (set.Set[int], float64) {
	if len(p) == 0 && len(x) == 0 {
		var weight float64
		for node := range r {
			weight += graph.GetWeight(node)
		}
		return r, weight
	}

	var maximalWeightClique set.Set[int]
	var maximalWeight float64
	for len(p) > 0 {
		v := p.Pop()
		vNeighbors := graph.GetNeighbors(v)
		rv := r.Clone().Add(v)
		pv := p.Clone().Intersect(vNeighbors)
		xv := x.Clone().Intersect(vNeighbors)
		clique, weight := BronKerbosch(rv, pv, xv, graph)
		x.Add(v)
		if weight > maximalWeight {
			maximalWeightClique = clique
			maximalWeight = weight
		}
	}

	return maximalWeightClique, maximalWeight
}
