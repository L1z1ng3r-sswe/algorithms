type Robot interface {
	Move() bool // is obstacle or not
	TurnLeft()
	TurnRight()
	Clean() // u can't clean the same area twice
}

type Bot interface {
	Robot
	back() bool
}

func (b bot) back(robot Robot) {
	b.TurnRight()
	b.robot.TurnRight()
	robot.Move()
	b.TurnRight()
	b.TurnRight()
}

// var dirs = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} // right, down, left, up

func CleanRoom(robot Robot) {
	bot := Bot(robot)

	visited := make(map[[2]int]bool)

	var backtrack func(x, y int)
	backtrack = func(x, y int) {
		pair := [2]int{x, y}

		if _, ok := visited[pair]; ok {
			bot.back()
			return
		}

		// move straight
		if bot.Move() {
			backtrack()
		}

		// move right
		bot.TurnRight()
		if bot.Move() {
			backtrack()
		}

		// move left
		bot.TurnRight()
		bot.TurnRight()
		if bot.Move() {
			backtrack()
		}

		// move back
		bot.TurnLeft()
		if bot.Move() {
			backtrack()
		}

		bot.Back()
	}

	backtrack()
}

