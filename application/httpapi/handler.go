package httpapi

import (
	"bytes"
	"github.com/allentom/haruka"
	"github.com/gabriel-vasile/mimetype"
	"github.com/projectxpolaris/thumbnailservice/service"
	"io"
	"net/http"
)

var generateThumbnailHandler haruka.RequestHandler = func(context *haruka.Context) {
	option := service.ThumbnailOption{}
	err := context.BindingInput(&option)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	r := context.Request
	r.ParseMultipartForm(32 << 32)
	file, _, err := r.FormFile("file") // Retrieve the file from form data
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	defer file.Close() // Close the file when we finish
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	thumbnail, err := service.MakeThumbnails(buf.Bytes(), option)
	mtype := mimetype.Detect(thumbnail)
	context.Writer.Header().Set("Content-Type", mtype.String())
	context.Writer.Write(thumbnail)
}

var infoHandler haruka.RequestHandler = func(context *haruka.Context) {
	context.JSON(haruka.JSON{
		"success": true,
	})
}
