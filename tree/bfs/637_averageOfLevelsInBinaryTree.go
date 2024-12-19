func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currLen := len(queue)
		sum := 0

		for i := range currLen {
			currNode := queue[i]
			sum += currNode.Val

			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}

		queue = queue[currLen:]
		res = append(res, float64(sum)/float64(currLen))
	}

	return res
}
