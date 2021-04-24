package rest

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	r.GET("/config/:appname", routesHandler.GetConfigHandler)
}
