package renderer

import (
	"image"
	"image/gif"
	"io"
	"time"
)

func MakeGif(w io.Writer, points []int, trackDuration time.Duration, conf Config) {
	delay := int(trackDuration.Seconds() / float64(len(points)) * 1000 / 10) // delay between frames (10ms)

	anim := gif.GIF{LoopCount: len(points)}

	points = Project(points, conf.Height)
	for _, v := range points {
		img := renderFrame(v, conf.Width, conf.Height, conf.ColorFunc)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(w, &anim)
}

func renderFrame(v int, width, height int, colorFunc ColorFunc) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, colorFunc.Palette())

	DrawRectangle(img, 0, height-v, width, height, height, colorFunc)

	return img
}
