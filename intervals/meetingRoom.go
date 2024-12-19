func canAttendMeeting(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { // in ASC order by a beginning time
		return intervals[i][0] < intervals[j][0]
	})

	lastEnd := intervals[0][1]

	for _, interval := range intervals[1:] {
		if lastEnd >= interval[0] {
			return false
		}
		lastEnd = interval[1]
	}

	return true
}