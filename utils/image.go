package utils

import (
	"image"
	"image/draw"
)

func ToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return rgba
}

func Crop(img image.Image, area image.Rectangle) *image.RGBA {
	base := ToRGBA(img)
	return ToRGBA(base.SubImage(area))
}
