package maze

import (
	"fmt"
	"math/rand"

	"kevyn.fr/maze/src"
)

type PrimsMaze struct {
	Maze
}

func (m *PrimsMaze) Generate() {
	start := rand.Intn(int(m.W * m.H))
	queue := []uint{uint(start)}

	order := 0

	for len(queue) > 0 {
		i := rand.Intn(len(queue))
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

	fmt.Println("\nPrim's maze generated.")
}
