package main

import (
	"github.com/julianstephens/graphae/utils"
	"github.com/op/go-logging"
)

func main() {
	logging.SetLevel(logging.DEBUG, utils.LOG.Module)
	_, err := OpenImage("img/bird.jpeg")
	if err != nil {
		utils.LOG.Fatal("something went wrong")
	}
}
