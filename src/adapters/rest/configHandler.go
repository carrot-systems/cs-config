package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) GetConfigHandler(c *gin.Context) {
	name := c.Param("appname")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "appname cannot be empty"})
	}

	config, err := rH.Usecases.GetConfig(name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}
