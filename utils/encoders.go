package utils

import (
	"image"
	"image/jpeg"
	"io"
)

type Encoder func(io.Writer, image.Image) error

func JPEGEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}
}
