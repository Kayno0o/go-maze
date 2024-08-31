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

type DrawParam struct {
	BlockSize     int
	SquarePadding int
	WallColor     color.Color
	WallWidth     int
	Rainbow       bool
}

func (m *Maze) Draw(filename string, params *DrawParam) {
	blockSize := params.BlockSize
	if blockSize == 0 {
		blockSize = 16
	}

	squarePadding := params.SquarePadding
	if squarePadding == 0 {
		squarePadding = 3
	}

	imgWidth := int(m.W) * blockSize
	imgHeight := int(m.H) * blockSize

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	wallColor := params.WallColor
	if wallColor == nil {
		wallColor = color.Black
	}

	wallWidth := params.WallWidth
	if wallWidth == 0 {
		wallWidth = 1
	}

	padding := wallWidth / 2

	src.DrawRect(img, 0, 0, imgWidth, imgHeight, color.White)

	// DRAW WALLS
	for _, block := range m.Blocks {
		x := int(block.Pos % m.W)
		y := int(block.Pos / m.W)

		// Calculate hue based on position
		if params.Rainbow == true {
			hue := float64((x + y) % 360)
			wallColor = src.HueToRGB(hue)
		}

		if block.Wall&1 == 0 {
			src.DrawRect(img, x*blockSize, y*blockSize-padding, blockSize, wallWidth, wallColor)
		}
		if block.Wall&2 == 0 {
			src.DrawRect(img, (x+1)*blockSize-padding, y*blockSize, wallWidth, blockSize, wallColor)
		}
		if block.Wall&4 == 0 {
			src.DrawRect(img, x*blockSize, (y+1)*blockSize-padding, blockSize, wallWidth, wallColor)
		}
		if block.Wall&8 == 0 {
			src.DrawRect(img, x*blockSize-padding, y*blockSize, wallWidth, blockSize, wallColor)
		}

		if block.Wall&1 == 0 && block.Wall&2 == 0 {
			src.DrawRect(img, (x+1)*blockSize-padding-min(1, padding/2), y*blockSize-padding+min(1, padding/2), wallWidth, wallWidth, wallColor)
		}
		if block.Wall&2 == 0 && block.Wall&4 == 0 {
			src.DrawRect(img, (x+1)*blockSize-padding-min(1, padding/2), (y+1)*blockSize-padding-min(1, padding/2), wallWidth, wallWidth, wallColor)
		}
		if block.Wall&4 == 0 && block.Wall&8 == 0 {
			src.DrawRect(img, x*blockSize-padding+min(1, padding/2), (y+1)*blockSize-padding-min(1, padding/2), wallWidth, wallWidth, wallColor)
		}
		if block.Wall&8 == 0 && block.Wall&1 == 0 {
			src.DrawRect(img, x*blockSize-padding+min(1, padding/2), y*blockSize-padding+min(1, padding/2), wallWidth, wallWidth, wallColor)
		}
	}

	// DRAW START/END
	startBlock := m.Blocks[rand.Intn(len(m.Blocks))]
	src.DrawRect(img, int(startBlock.Pos%m.W)*blockSize+squarePadding, int(startBlock.Pos/m.W)*blockSize+squarePadding, blockSize-squarePadding*2, blockSize-squarePadding*2, color.RGBA{0, 200, 255, 255})

	endBlock := m.Blocks[rand.Intn(len(m.Blocks))]
	src.DrawRect(img, int(endBlock.Pos%m.W)*blockSize+squarePadding, int(endBlock.Pos/m.W)*blockSize+squarePadding, blockSize-squarePadding*2, blockSize-squarePadding*2, color.RGBA{255, 200, 0, 255})

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
