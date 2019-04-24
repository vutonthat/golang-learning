package server

import (
	"digin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const PublicIPPath = "/public-ip"

type publicIPHandler struct {
	ginEngine       *gin.Engine
	publicIPService services.PublicIPService
}

func (_this *publicIPHandler) router() {
	_this.ginEngine.GET(PublicIPPath, _this.handleGet)
}

func (_this *publicIPHandler) handleGet(c *gin.Context) {
	resp := _this.publicIPService.GetPublicIP()
	c.JSON(http.StatusOK, resp)
}
