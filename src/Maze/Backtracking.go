package maze

import (
	"fmt"
	"math/rand"

	"kevyn.fr/maze/src"
)

type BacktrackingMaze struct {
	Maze
}

func (m *BacktrackingMaze) Generate() {
	start := rand.Intn(int(m.W * m.H))
	queue := []uint{uint(start)}

	order := 0

	for len(queue) > 0 {
		pos := queue[0]

		newPos, dir, err := m.RandomNonVisitedNeighbour(pos)
		if err != nil {
			queue = queue[1:]
			continue
		}

		m.Blocks[pos].Wall |= dir
		m.Blocks[newPos].Wall |= src.ReverseWall(dir)

		m.Blocks[newPos].Order = uint(order)
		order++

		queue = append([]uint{newPos}, queue...)
	}

	fmt.Println("\nRecursive backtracking maze generated.")
}
