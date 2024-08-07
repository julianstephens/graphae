package main

import (
	"bufio"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/julianstephens/graphae/utils"
)

type Image struct {
	im     image.Image
	format string
}

func OpenImage(filepath string) (*Image, error) {
	imgFile, err := os.Open(filepath)
	err = utils.HandleError(err, "unable to read image file")
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, format, err := image.Decode(bufio.NewReader(imgFile))
	err = utils.HandleError(err, "unable to decode image")
	if err != nil {
		return nil, err
	}

	utils.LOG.Debugf("decoded %s", format)

	return &Image{im: img, format: format}, nil
}

// func (i *Image) buffer() image.Image {
// 	return i.im
// }

// func (i *Image) crop(area []int) {}

// func (i *Image) colors(area []int) {}

// func (i *Image) size() {}
