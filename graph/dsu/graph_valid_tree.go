package algorithms

func ifValidTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	dsu := NewDSU(n)

	for _, edge := range edges {
		if !dsu.Union(edge[0], edge[1]) {
			return false
		}
	}

	return true
}

type DSU struct {
	Parent []int
	Rank   []int
}

func NewDSU(n int) *DSU {
	var dsu = &DSU{
		Parent: make([]int, n),
		Rank:   make([]int, n),
	}

	for i := range dsu.Parent {
		dsu.Parent[i] = i
	}
	return dsu
}

func (d *DSU) Find(x int) int {
	if d.Parent[x] != x {
		d.Parent[x] = d.Find(d.Parent[x])
	}

	return d.Parent[x]
}

// false only if a and b got the same Parent root (cycle detected)
func (d *DSU) Union(a, b int) bool {
	rootA, rootB := d.Find(a), d.Find(b)

	if rootA == rootB { // cycle detected
		return false
	}

	if d.Rank[rootA] < d.Rank[rootB] {
		d.Parent[rootA] = rootB
	} else if d.Rank[rootA] > d.Rank[rootB] {
		d.Parent[rootB] = rootA
	} else {
		d.Rank[rootA]++
		d.Parent[rootB] = rootA
	}

	return true
}

// m - len(edges)
// n - number of nodes

// Time:  O(m * α(n))   // ~ O(m)
// Space: O(n)
