package src

import (
	"image"
	"image/color"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func DrawLine(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	dx := Abs(x2 - x1)
	dy := Abs(y2 - y1)
	sx := -1
	if x1 < x2 {
		sx = 1
	}
	sy := -1
	if y1 < y2 {
		sy = 1
	}
	err := dx - dy

	for {
		img.Set(x1, y1, c)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func DrawRect(img *image.RGBA, x, y, w, h int, c color.Color) {
	for Y := y; Y < h; Y++ {
		for X := x; X < w; X++ {
			img.Set(X, Y, c)
		}
	}
}
