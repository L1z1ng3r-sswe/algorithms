var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	visited := make([]bool, m*n) // Flattened 2D array to 1D for better cache
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, &Item{0, 0, 0, 0})

	for {
		item := heap.Pop(minHeap).(*Item)
		if visited[item.idx] {
			continue
		}
		visited[item.idx] = true

		if item.r == m-1 && item.c == n-1 {
			return item.diff
		}

		for _, dir := range dirs {
			newR, newC := item.r+dir[0], item.c+dir[1]
			newIdx := newR*n + newC

			if newR < 0 || newC < 0 || newR >= m || newC >= n || visited[newIdx] {
				continue
			}

			newDiff := heights[newR][newC] - heights[item.r][item.c]
			if newDiff < 0 {
				newDiff = -newDiff
			}

			if newDiff < item.diff {
				newDiff = item.diff
			}

			heap.Push(minHeap, &Item{newR, newC, newDiff, newIdx})
		}
	}
}

type MinHeap []*Item

type Item struct {
	r, c int
	diff int
	idx  int // Flattened idx not to recalculate
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i].diff < h[j].diff }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}