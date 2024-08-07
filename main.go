package main

import (
	"image"

	"github.com/julianstephens/graphae/utils"
	"github.com/op/go-logging"
)

func main() {
	logging.SetLevel(logging.DEBUG, utils.LOG.Module)
	img, err := utils.Open("img/bird.jpeg")
	err = utils.HandleError(err, "something went wrong")
	if err != nil {
		utils.LOG.Panic("something went wrong")
	}

	cropped := utils.Crop(img, image.Rect(0, 0, 100, 100))
	err = utils.Write("img/res.jpeg", cropped, utils.JPEGEncoder(100))
	if err != nil {
		utils.LOG.Panic("something went wrong")
	}
}
