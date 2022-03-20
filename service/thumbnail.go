package service

import (
	"github.com/h2non/bimg"
)

type ThumbnailOption struct {
	MaxWidth  int    `hsource:"query" hname:"maxWidth"`
	MaxHeight int    `hsource:"query" hname:"maxHeight"`
	Mode      string `hsource:"query" hname:"mode"`
}

func MakeThumbnails(input []byte, option ThumbnailOption) ([]byte, error) {
	sourceImage := bimg.NewImage(input)
	size, err := sourceImage.Size()
	if err != nil {
		return nil, err
	}
	thumbnailWidth := 0
	thumbnailHeight := 0
	switch option.Mode {
	case "width":
		thumbnailWidth = option.MaxWidth
		thumbnailHeight = int(float64(option.MaxWidth) * float64(size.Height) / float64(size.Width))
	case "height":
		thumbnailHeight = option.MaxHeight
		thumbnailWidth = int(float64(option.MaxHeight) * float64(size.Width) / float64(size.Height))
	case "resize":
		thumbnailWidth = option.MaxWidth
		thumbnailHeight = option.MaxHeight
	default:
		widthRatio := float64(size.Width) / float64(option.MaxWidth)
		heightRatio := float64(size.Height) / float64(option.MaxHeight)
		if widthRatio > heightRatio {
			thumbnailWidth = option.MaxWidth
			thumbnailHeight = int(float64(option.MaxWidth) * float64(size.Height) / float64(size.Width))
		} else {
			thumbnailHeight = option.MaxHeight
			thumbnailWidth = int(float64(option.MaxHeight) * float64(size.Width) / float64(size.Height))
		}
	}
	newImage, err := sourceImage.Resize(thumbnailWidth, thumbnailHeight)
	if err != nil {
		return nil, err
	}
	return newImage, nil
}
