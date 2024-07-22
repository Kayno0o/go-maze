package src

// 1 = top, 2 = right, 4 = bottom, 8 = left
var Dirs = []uint8{1, 2, 4, 8}

type Block struct {
	Pos   uint
	Order uint
	Wall  uint8
}

func ReverseWall(wall uint8) uint8 {
	switch wall {
	case 1:
		return 4
	case 2:
		return 8
	case 4:
		return 1
	case 8:
		return 2
	}
	return 0
}
