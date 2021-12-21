package main

import (
	"github.com/NETkiddy/nft-svr/common/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/NETkiddy/common-go/svr_adapter/glue"
	"github.com/NETkiddy/common-go/svr_adapter/glue/middlewares/gin_middleware"
	"github.com/NETkiddy/nft-svr/token/handler"
)

func InitRoute() http.Handler {
	Routes := glue.Gin()

	logMiddleware := gin_middleware.NewLogMiddlewareWithCfgSystem(nil)
	Routes.Use(logMiddleware.Log())
	Routes.Use(gin.Recovery())

	vProduct := Routes.Group("nft/token_classes").Use(middlewares.CheckToken())
	{
		vProduct.POST("", glue.UseControllerWithContextForCGIAPIV3(handler.NewToken().TokenClasses))
		vProduct.POST("Delete", glue.UseControllerWithContextForCGIAPIV3(handler.NewToken().TokenClasses))
	}

	return Routes
}
