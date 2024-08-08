package cli

import (
	"fmt"
	"image"

	"github.com/julianstephens/graphae/utils"
)

type CropCmd struct {
	SrcImagePath  string `help:"The path to the image to crop" short:"s"`
	DestImagePath string `help:"The path to save the cropped image" short:"d"`
	CropArea      []int  `help:"The 4 bounds of the crop area" short:"a"`
}

type TransformCmd struct {
	Crop CropCmd `cmd:"" help:"Crop an image to the specified area"`
}

func (c *CropCmd) Run(globals *Globals) error {
	if len(c.CropArea) != 4 {
		panic("invalid crop area. must be 4 valid integers")
	}

	img, err := utils.Open(c.SrcImagePath)
	if err != nil {
		panic(fmt.Sprintf("unable to open image %s", c.SrcImagePath))
	}

	cropped := utils.Crop(img, image.Rect(c.CropArea[0], c.CropArea[1], c.CropArea[2], c.CropArea[3]))

	err = utils.Write(c.DestImagePath, cropped, utils.JPEGEncoder(100))
	if err != nil {
		panic(err)
	}

	return nil
}
