package maze

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"kevyn.fr/maze/src"
)

type Maze struct {
	W      uint
	H      uint
	Blocks []*src.Block
}

func (m *Maze) Init(w, h uint) {
	m.W = w
	m.H = h
	m.Blocks = make([]*src.Block, w*h)

	for i := 0; i < int(w*h); i++ {
		m.Blocks[i] = &src.Block{Pos: uint(i), Wall: 0, Order: 0}
	}
}

func (m *Maze) GetBlock(pos uint, dir uint8) (uint, error) {
	newPos := -1

	switch dir {
	case 1:
		if pos >= m.W {
			newPos = int(pos - m.W)
		}
	case 2:
		if (pos+1)%m.W != 0 {
			newPos = int(pos + 1)
		}
	case 4:
		if pos < m.W*(m.H-1) {
			newPos = int(pos + m.W)
		}
	case 8:
		if pos%m.W != 0 {
			newPos = int(pos - 1)
		}
	}

	if newPos < 0 || newPos >= int(m.W*m.H) {
		return 0, fmt.Errorf("out of range")
	}

	return uint(newPos), nil
}

func (m *Maze) GetNeighbours(pos uint) []uint {
	neighbours := []uint{}

	for _, dir := range src.Dirs {
		newPos, err := m.GetBlock(pos, dir)
		if err != nil {
			continue
		}

		neighbours = append(neighbours, newPos)
	}

	return neighbours
}

func (m *Maze) GetNonVisitedNeighbours(pos uint) []uint {
	nonVisitedNeighbours := []uint{}
	neighbours := m.GetNeighbours(pos)

	for _, neighbour := range neighbours {
		if m.Blocks[neighbour].Wall != 0 {
			continue
		}

		nonVisitedNeighbours = append(nonVisitedNeighbours, neighbour)
	}

	return nonVisitedNeighbours
}

func (m *Maze) RandomNonVisitedNeighbour(pos uint) (uint, uint8, error) {
	shuffledDirs := append([]uint8(nil), src.Dirs...)
	rand.Shuffle(len(shuffledDirs), func(i, j int) { shuffledDirs[i], shuffledDirs[j] = shuffledDirs[j], shuffledDirs[i] })

	for _, dir := range shuffledDirs {
		newPos, err := m.GetBlock(pos, dir)
		if err != nil {
			continue
		}

		if m.Blocks[newPos].Wall == 0 {
			return newPos, dir, nil
		}
	}

	return 0, 0, fmt.Errorf("no neighbours")
}

func (m *Maze) Draw(filename string) {
	blockSize := 25
	padding := 3

	imgWidth := int(m.W) * blockSize
	imgHeight := int(m.H) * blockSize

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	wallColor := color.RGBA{0, 0, 0, 255}

	src.DrawRect(img, 0, 0, imgWidth, imgHeight, color.White)

	// DRAW WALLS
	for _, block := range m.Blocks {
		x := int(block.Pos % m.W)
		y := int(block.Pos / m.W)

		if block.Wall&1 == 0 {
			src.DrawLine(img, x*blockSize, y*blockSize, (x+1)*blockSize, y*blockSize, wallColor)
		}
		if block.Wall&2 == 0 {
			src.DrawLine(img, (x+1)*blockSize, y*blockSize, (x+1)*blockSize, (y+1)*blockSize, wallColor)
		}
		if block.Wall&4 == 0 {
			src.DrawLine(img, x*blockSize, (y+1)*blockSize, (x+1)*blockSize, (y+1)*blockSize, wallColor)
		}
		if block.Wall&8 == 0 {
			src.DrawLine(img, x*blockSize, y*blockSize, x*blockSize, (y+1)*blockSize, wallColor)
		}
	}

	// DRAW START/END
	drawCenter := func(block *src.Block, c color.Color) {
		x := int(block.Pos % m.W)
		y := int(block.Pos / m.W)
		startX := x*blockSize + padding
		startY := y*blockSize + padding
		endX := (x+1)*blockSize - padding
		endY := (y+1)*blockSize - padding

		for i := startX; i < endX; i++ {
			for j := startY; j < endY; j++ {
				img.Set(i, j, c)
			}
		}
	}

	drawCenter(m.Blocks[rand.Intn(len(m.Blocks))], color.RGBA{0, 200, 255, 255})
	drawCenter(m.Blocks[rand.Intn(len(m.Blocks))], color.RGBA{255, 200, 0, 255})

	// SAVE IMAGE
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)

	if err != nil {
		panic(err)
	}

	fmt.Println("Maze saved as " + filename + ".")
}
