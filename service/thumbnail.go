package service

import (
	"github.com/davidbyttow/govips/v2/vips"
)

type ThumbnailOption struct {
	MaxWidth  int    `hsource:"query" hname:"maxWidth"`
	MaxHeight int    `hsource:"query" hname:"maxHeight"`
	Mode      string `hsource:"query" hname:"mode"`
}

func MakeThumbnails(input []byte, option ThumbnailOption) ([]byte, error) {
	sourceImage, err := vips.NewImageFromBuffer(input)
	if err != nil {
		return nil, err
	}
	sourceImage.Width()
	thumbnailWidth := 0
	thumbnailHeight := 0
	switch option.Mode {
	case "width":
		thumbnailWidth = option.MaxWidth
		thumbnailHeight = int(float64(option.MaxWidth) * float64(sourceImage.Height()) / float64(sourceImage.Width()))
	case "height":
		thumbnailHeight = option.MaxHeight
		thumbnailWidth = int(float64(option.MaxHeight) * float64(sourceImage.Width()) / float64(sourceImage.Height()))
	case "resize":
		thumbnailWidth = option.MaxWidth
		thumbnailHeight = option.MaxHeight
	default:
		widthRatio := float64(sourceImage.Width()) / float64(option.MaxWidth)
		heightRatio := float64(sourceImage.Height()) / float64(option.MaxHeight)
		if widthRatio > heightRatio {
			thumbnailWidth = option.MaxWidth
			thumbnailHeight = int(float64(option.MaxWidth) * float64(sourceImage.Height()) / float64(sourceImage.Width()))
		} else {
			thumbnailHeight = option.MaxHeight
			thumbnailWidth = int(float64(option.MaxHeight) * float64(sourceImage.Width()) / float64(sourceImage.Height()))
		}
	}
	err = sourceImage.Thumbnail(thumbnailWidth, thumbnailHeight, vips.InterestingAll)
	if err != nil {
		return nil, err
	}
	image1bytes, _, err := sourceImage.ExportPng(vips.NewPngExportParams())
	if err != nil {
		return nil, err
	}

	return image1bytes, nil
}
