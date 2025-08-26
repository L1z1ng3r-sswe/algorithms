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
	Size   []int
}

func NewDSU(n int) *DSU {
	var dsu = &DSU{
		Parent: make([]int, n),
		Size:   make([]int, n),
	}

	for i := range dsu.Parent {
		dsu.Parent[i] = i
		dsu.Size[i] = 1
	}
	return dsu
}

// Find uses iterative find and path compression
func (d *DSU) Find(x int) int {
	root := x
	for d.Parent[root] != root {
		root = d.Parent[root]
	}

	// path compression
	for d.Parent[x] != x {
		nxt := d.Parent[x]
		d.Parent[x] = root
		x = nxt
	}

	return root
}

// false only if a and b got the same Parent root (cycle detected)
func (d *DSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)

	if ra == rb { // cycle detected
		return false
	}

	if d.Size[ra] < d.Size[rb] {
		d.Parent[ra] = rb
		d.Size[rb] += d.Size[ra]
	} else {
		d.Parent[rb] = ra
		d.Size[ra] += d.Size[rb]
	}

	return true
}

// m - len(edges)
// n - number of nodes

// Time:  O(m * α(n))   // ~ O(m)
// Space: O(n)
