package main

import (
	"github.com/NETkiddy/nft-svr/common/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/NETkiddy/common-go/svr_adapter/glue"
	"github.com/NETkiddy/common-go/svr_adapter/glue/middlewares/gin_middleware"
	"github.com/NETkiddy/nft-svr/provider/handler"
)

func InitRoute() http.Handler {
	Routes := glue.Gin()

	logMiddleware := gin_middleware.NewLogMiddlewareWithCfgSystem(nil)
	Routes.Use(logMiddleware.Log())
	Routes.Use(gin.Recovery())

	vArtist := Routes.Group("bee/provider").Use(middlewares.CheckToken())
	{
		//vArtist.POST("Create", glue.UseControllerWithContextForCGIAPIV3(handler.NewProvider().CreateProvider))
		//vArtist.POST("Delete", glue.UseControllerWithContextForCGIAPIV3(handler.NewProvider().DeleteProvider))
		vArtist.POST("Query", glue.UseControllerWithContextForCGIAPIV3(handler.NewProvider().QueryProvider))
		//vArtist.POST("Update", glue.UseControllerWithContextForCGIAPIV3(handler.NewProvider().UpdateProvider))
	}

	return Routes
}
