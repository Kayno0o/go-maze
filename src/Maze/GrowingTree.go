package maze

import (
	"fmt"
	"math/rand"

	"kevyn.fr/maze/src"
)

type GrowingTreeMazeInterface interface {
	GetCell(queue []uint) string
}

type GrowingTreeMaze struct {
	Maze
}

func (m *GrowingTreeMaze) Generate(GetCell func([]uint) int) {
	start := rand.Intn(int(m.W * m.H))
	queue := []uint{uint(start)}

	order := 0

	for len(queue) > 0 {
		i := GetCell(queue)
		pos := queue[i]

		newPos, dir, err := m.RandomNonVisitedNeighbour(pos)
		if err != nil {
			queue = append(queue[0:i], queue[i+1:]...)
			continue
		}

		m.Blocks[pos].Wall |= dir
		m.Blocks[newPos].Wall |= src.ReverseWall(dir)

		m.Blocks[newPos].Order = uint(order)
		order++

		queue = append([]uint{newPos}, queue...)
	}

	fmt.Println("\nGrowing tree maze generated.")
}
