package rest

import (
	"config/src/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) GetConfigHandler(c *gin.Context) {
	name := c.Param("appname")

	if name == "" {
		c.JSON(http.StatusBadRequest, domain.ConfigResponse{
			Status: domain.Status{
				Success: false,
				Message: "appname parameter cannot be empty",
			},
			Values: nil,
		})
	}

	configValues, err := rH.Usecases.GetConfig(name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, domain.ConfigResponse{
			Status: domain.Status{
				Success: false,
				Message: err.Error(),
			},
			Values: nil,
		})
		return
	}

	c.JSON(http.StatusOK, domain.ConfigResponse{
		Status: domain.Status{
			Success: true,
			Message: "",
		},
		Values: configValues,
	})
}
