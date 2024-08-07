package main

import (
	"testing"
)

func TestNewImageFilename(t *testing.T) {
	filename := "img/bird.jpeg"

	img, err := OpenImage(filename)
	if err != nil {
		t.Fatalf("could not create new image...%v", err)
	}

	if img.format != "jpeg" {
		t.Fatalf("unknown image format...%s", img.format)
	}
}
