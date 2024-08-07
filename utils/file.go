package utils

import (
	"bufio"
	"image"
	"os"
)

func Open(filepath string) (image.Image, error) {
	imgFile, err := os.Open(filepath)
	err = HandleError(err, "unable to read image file")
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, format, err := image.Decode(bufio.NewReader(imgFile))
	err = HandleError(err, "unable to decode image")
	if err != nil {
		return nil, err
	}

	LOG.Debugf("decoded %s", format)

	return img, nil
}

func Write(filepath string, img image.Image, encoder Encoder) error {
	f, err := os.Create(filepath)
	err = HandleError(err, "could not create image file")
	if err != nil {
		LOG.Panic("something went wrong")
	}
	defer f.Close()

	return encoder(f, img)
}
