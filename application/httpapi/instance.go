package httpapi

import (
	"github.com/allentom/haruka"
	"github.com/allentom/haruka/middleware"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var Logger = log.New().WithFields(log.Fields{
	"scope": "Application",
})

func GetEngine() *haruka.Engine {
	e := haruka.NewEngine()
	e.UseCors(cors.AllowAll())
	e.UseMiddleware(middleware.NewLoggerMiddleware())
	e.UseMiddleware(middleware.NewPaginationMiddleware("page", "pageSize", 1, 20))
	e.Router.POST("/generator", generateThumbnailHandler)
	e.Router.GET("/info", infoHandler)
	return e
}
