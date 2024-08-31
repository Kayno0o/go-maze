package src

import (
	"image"
	"image/color"
	"math"
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
	for Y := y; Y < y+h; Y++ {
		for X := x; X < x+w; X++ {
			img.Set(X, Y, c)
		}
	}
}

func HueToRGB(hue float64) color.RGBA {
	h := hue / 60
	x := 1 - math.Abs(math.Mod(h, 2)-1)

	var r, g, b float64
	switch {
	case h >= 0 && h < 1:
		r, g, b = 1, x, 0
	case h >= 1 && h < 2:
		r, g, b = x, 1, 0
	case h >= 2 && h < 3:
		r, g, b = 0, 1, x
	case h >= 3 && h < 4:
		r, g, b = 0, x, 1
	case h >= 4 && h < 5:
		r, g, b = x, 0, 1
	case h >= 5 && h < 6:
		r, g, b = 1, 0, x
	}

	return color.RGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: 255,
	}
}
