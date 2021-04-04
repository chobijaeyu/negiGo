package routers

import (
	"negigo/middleware"
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

	r.Use(middleware.FirebaseAuth())

	r.Use(middleware.OperatingLog())

	r.POST("v1/img/:name", views.AddGoodsImg)
	r.DELETE("v1/img", views.DeleteGoodsImg)

	var authViews views.AuthViews
	// r.GET("/v1/members", authViews.ListAllUsers)
	// r.POST("/v1/members/:uid", authViews.UpdateUser)
	r.GET("/v1/members", authViews.ListAllUsers)
	r.POST("/v1/members/:uid", middleware.AdminAuth(), authViews.UpdateUser)

	var fieldview views.NegiField
	r.GET("v1/negifields/", fieldview.GetAllNeigFields)
	fieldRouterGroup := r.Group("/v1/negifield/")
	{
		fieldRouterGroup.GET(":negifieldid", fieldview.GetNegiField)
		fieldRouterGroup.POST("", fieldview.CreateNegiField)
		fieldRouterGroup.PUT(":negifieldid", fieldview.UpdateNegiField)
		fieldRouterGroup.DELETE(":negifieldid", fieldview.DeteleNegiField)
	}

	var taskcaleventview views.TaskCalEvent
	r.GET("v1/negicalevents/", taskcaleventview.GetAllTaskCalEvent)
	taskRouterGroup := r.Group("/v1/negicalevent/")
	{
		taskRouterGroup.GET("")
		taskRouterGroup.POST("", taskcaleventview.CreateTaskCalEvent)
		taskRouterGroup.PUT(":eventid", taskcaleventview.UpdateTaskCalEvent)
		taskRouterGroup.DELETE(":eventid", taskcaleventview.DeteleTaskCalEvent)
	}

	var tasktitleoptionview views.NegiTaskTitleOption
	r.GET("v1/negicustomtasktitleoptions/", tasktitleoptionview.GetAllTaskTitleOption)
	tasktitleoptionRouterGroup := r.Group("/v1/negicustomtasktitleoption/")
	{
		tasktitleoptionRouterGroup.GET("")
		tasktitleoptionRouterGroup.POST("", tasktitleoptionview.CreateTaskTitleOption)
		tasktitleoptionRouterGroup.PUT(":id", tasktitleoptionview.UpdateTaskTitleOption)
		tasktitleoptionRouterGroup.DELETE(":id", tasktitleoptionview.DeleteTaskTitleOption)
	}

	var seriestaskoptionview views.NegiSeriesTaskOption
	r.GET("v1/negicustomseriestaskoptions/", seriestaskoptionview.GetAllseriesTaskOption)
	seriestaskoptionRouterGroup := r.Group("/v1/negicustomseriestaskoption/")
	{
		seriestaskoptionRouterGroup.GET("")
		seriestaskoptionRouterGroup.POST("", seriestaskoptionview.CreateSeriesTaskOption)
		seriestaskoptionRouterGroup.PUT(":id", seriestaskoptionview.UpdateseriesTaskOption)
		seriestaskoptionRouterGroup.DELETE(":id", seriestaskoptionview.DeleteseriesTaskOption)
	}
	return r
}
