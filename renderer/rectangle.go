package renderer

import (
	"image/color"
)

type Img interface {
	Set(x, y int, c color.Color)
}

// Draw Rectangle receives an image and draws a rectangle using colorFunc
// iterate over x and y axis:
// x: from x1 to x2
// y: from y1 to y2
// Use colorFunc.Color(x,y,height) to determine the color to paint each pixel
// To color a pixel use img.Set

func DrawRectangle(img Img, x1, y1, x2, y2 int, height int, colorFunc ColorFunc) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			c := colorFunc.Color(x, x1, y, y1, height)
			img.Set(x, y, c)
		}
	}
}
