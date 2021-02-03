package routers

import (
	"negigo/views"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "Access-Control-Allow-Origin", "content-type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
		// AllowOrigins:     []string{"http://10.*.*.*:*", "http://localhost:*", "http://127.0.0.1:*", "http://192.168.50.*:*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
		// AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 12 * time.Hour,

		// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
		AllowWildcard: true,

		// Allows usage of popular browser extensions schemas
		AllowBrowserExtensions: true,

		// Allows usage of WebSocket protocol
		AllowWebSockets: true,

		// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
		AllowFiles: true,
	}))

	var fieldview views.NegiField
	r.GET("v1/negifields/", fieldview.GetAllNeigFields)
	fieldRouterGroup := r.Group("/v1/negifield/")
	{
		fieldRouterGroup.GET(":negifieldid", fieldview.GetNegiField)
		fieldRouterGroup.POST("", fieldview.CreateNegiField)
		fieldRouterGroup.PATCH("", fieldview.UpdateNegiField)
		fieldRouterGroup.DELETE("", fieldview.DeteleNegiField)
	}

	var taskcaleventview views.TaskCalEvent
	taskRouterGroup := r.Group("/v1/negicalevent/")
	{
		taskRouterGroup.GET("")
		taskRouterGroup.POST("", taskcaleventview.CreateTaskCalEvent)
		taskRouterGroup.PATCH("", taskcaleventview.UpdateTaskCalEvent)
		taskRouterGroup.DELETE("", taskcaleventview.DeteleTaskCalEvent)
	}

	return r
}
