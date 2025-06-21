type Robot interface {
	Move() bool // is obstacle or not
	TurnLeft()
	TurnRight()
	Clean() // u can't clean the same area twice
}

// up, right, down, left
var dirs = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up is the first because the initial direction is up

func CleanRoom(robot Robot) {
	var back = func() {
		robot.TurnRight()
		robot.TurnRight()
		robot.Move()
		robot.TurnRight()
		robot.TurnRight()
	}
	visited := make(map[[2]int]struct{})

	var backtrack func(x, y int, dir int)
	backtrack = func(x, y int, dir int) {
		visited[[2]int{x, y}] = struct{}{}
		robot.Clean()

		for i := 0; i < 4; i++ {
			newDir := (dir + i) % 4
			newX := x + dirs[newDir][0]
			newY := y + dirs[newDir][1]

			if _, ok := visited[[2]int{newX, newY}]; !ok {
				if robot.Move() {
					backtrack(newX, newY, newDir)
					back()
				}
			}
			robot.TurnRight()
		}
	}

	backtrack(0, 0, 0)
}