package httpapi

import (
	"github.com/allentom/haruka"
	"github.com/projectxpolaris/thumbnailservice/commons"
	"github.com/projectxpolaris/thumbnailservice/youlog"
)

func AbortError(ctx *haruka.Context, err error, status int) {
	if apiError, ok := err.(*commons.APIError); ok {
		youlog.DefaultYouLogPlugin.Logger.Error(apiError.Err.Error())
		ctx.JSONWithStatus(haruka.JSON{
			"success": false,
			"err":     apiError.Desc,
			"code":    apiError.Code,
		}, status)
		return
	}
	youlog.DefaultYouLogPlugin.Logger.Error(err.Error())
	ctx.JSONWithStatus(haruka.JSON{
		"success": false,
		"err":     err.(error).Error(),
		"code":    "9999",
	}, status)
}
