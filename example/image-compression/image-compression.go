package imagecompression

import (
	"image"

	"github.com/disintegration/imaging"
)

func CompressImage(img image.Image, width, height int) image.Image {
	dstImage := imaging.Resize(img, width, height, imaging.Lanczos)
	return dstImage
}
